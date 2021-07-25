package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/felixa1996/go-plate/infrastructure/route"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/felixa1996/go-plate/adapter/api/action"
	"github.com/felixa1996/go-plate/adapter/api/middleware"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/adapter/presenter"
	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/adapter/validator"
	"github.com/felixa1996/go-plate/usecase"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	_ "github.com/felixa1996/go-plate/docs"
	//openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	db         repository.SQL
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	db repository.SQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *gorillaMux {
	return &gorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		db:         db,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

// @title Data API
// @version 1.0
// @description Data API
// @termsOfService http://swagger.io/terms/
// @contact.email felixanthony1996.fa@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000/vhry/data/
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (g gorillaMux) Listen() {
	g.setAppHandlers(g.router)
	g.middleware.UseHandler(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.middleware,
	}

	g.router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			g.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	g.log.Infof("Service down")
}

func (g gorillaMux) setAppHandlers(router *mux.Router) {
	api := router.PathPrefix("/v1").Subrouter()

	api.Handle("/transfers", g.buildFindAllTransferAction()).Methods(http.MethodGet)

	api.Handle("/accounts/{account_id}/balance", g.buildFindBalanceAccountAction()).Methods(http.MethodGet)
	api.Handle("/accounts/{id}", g.buildDeleteBalanceAccountAction()).Methods(http.MethodDelete)
	api.Handle("/accounts", g.buildCreateAccountAction()).Methods(http.MethodPost)
	api.Handle("/accounts", g.buildFindAllAccountAction()).Methods(http.MethodGet)

	api.Handle("/charity-mrys", g.buildCreateCharityMrysAction()).Methods(http.MethodPost)
	api.Handle("/charity-mrys/{id}", g.buildUpdateCharityMrysAction()).Methods(http.MethodPatch)
	api.Handle("/charity-mrys/create-bulk", g.buildCreateBulkCharityMrysAction()).Methods(http.MethodPost)
	api.Handle("/charity-mrys/list-pagination/{currentPage}/{perPage}/{sort}", g.buildFindPaginationCharityMrysAction()).
		Queries("search", "{search}").
		Methods(http.MethodGet)
	api.Handle("/charity-mrys/list-all", g.buildFindAllCharityMrysAction()).Methods(http.MethodGet)
	api.Handle("/charity-mrys/{id}", g.buildFindCharityMrysAction()).Methods(http.MethodGet)
	api.Handle("/charity-mrys/{id}", g.buildDeleteOneCharityMrysAction()).Methods(http.MethodDelete)

	api.HandleFunc("/health", action.HealthCheck).Methods(http.MethodGet)
}

func (g gorillaMux) buildFindAllTransferAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewFindAllTransferInteractor(
				repository.NewTransferSQL(g.db),
				presenter.NewFindAllTransferPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllTransferAction(uc, g.log)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// CreateAccount godoc
// @Summary Create account
// @Description Create account
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param data body domain.Account true "Create account"
// @Success 201 {object} domain.Account
// @Router /v1/accounts [post]
func (g gorillaMux) buildCreateAccountAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewCreateAccountInteractor(
				repository.NewAccountSQL(g.db),
				presenter.NewCreateAccountPresenter(),
				g.ctxTimeout,
			)
			act = action.NewCreateAccountAction(uc, g.log, g.validator)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// FindAccounts godoc
// @Summary Find All Accounts
// @Tags Accounts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Account
// @Router /v1/accounts [get]
func (g gorillaMux) buildFindAllAccountAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.AccountFindAll(g.db, g.log, g.ctxTimeout)
		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// FindCharityMrys godoc
// @Summary Find All CharityMrys
// @Tags CharityMrys
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.CharityMrys
// @Router /v1/charity-mrys/list-all [get]
func (g gorillaMux) buildFindAllCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysFindPagination(g.db, g.log, g.ctxTimeout)
		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// FindPaginationCharityMrys godoc
// @Summary Find Pagination CharityMrys
// @Tags CharityMrys
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.CharityMrysPagination
// @Param currentPage path int true "CurrentPage"
// @Param perPage path int true "PerPage"
// @Param sort path string true "Sort"
// @Param search query string false "Search"
// @Router /v1/charity-mrys/list-pagination/{currentPage}/{perPage}/{sort} [get]
func (g gorillaMux) buildFindPaginationCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysFindPagination(g.db, g.log, g.ctxTimeout)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("currentPage", vars["currentPage"])
		q.Add("perPage", vars["perPage"])
		q.Add("sort", vars["sort"])
		q.Add("search", vars["search"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// CreateCharityMrys godoc
// @Summary Create Charity Mrys
// @Description Create Charity Mrys
// @Tags CharityMrys
// @Accept  json
// @Produce  json
// @Param data body domain.CharityMrys true "Create charity mrys"
// @Success 201 {object} domain.CharityMrys
// @Router /v1/charity-mrys [post]
func (g gorillaMux) buildCreateCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysCreateOne(g.db, g.log, g.ctxTimeout, g.validator)
		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// UpdateCharityMrys godoc
// @Summary Update Charity Mrys By ID
// @Description Update Charity Mrys By ID
// @Tags CharityMrys
// @Accept  json
// @Produce  json
// @Param data body domain.CharityMrys true "Update charity mrys"
// @Success 201 {object} domain.CharityMrys
// @Param id path string true "ID"
// @Router /v1/charity-mrys/{id} [patch]
func (g gorillaMux) buildUpdateCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysUpdateOne(g.db, g.log, g.ctxTimeout, g.validator)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("id", vars["id"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// CreateBulkCharityMrys godoc
// @Summary Create Bulk Charity Mrys
// @Description Create Bulk Charity Mrys
// @Tags CharityMrys
// @Accept json
// @Produce  json
// @Param data body usecase.CreateBulkCharityMrysInput true "Create charity mrys"
// @Success 201 {object} []domain.CharityMrys
// @Router /v1/charity-mrys/create-bulk [post]
func (g gorillaMux) buildCreateBulkCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysCreateBulk(g.db, g.log, g.ctxTimeout, g.validator)
		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// FindCharityMrys godoc
// @Summary Find One Charity Mrys By ID
// @Tags CharityMrys
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.CharityMrys
// @Param id path string true "ID"
// @Router /v1/charity-mrys/{id} [get]
func (g gorillaMux) buildFindCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysFindOne(g.db, g.log, g.ctxTimeout)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("id", vars["id"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// DeleteCharityMrys godoc
// @Summary Delete One Charity Mrys By ID
// @Tags CharityMrys
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {boolean} boolean "success"
// @Param id path string true "Charity Mrys ID"
// @Router /v1/charity-mrys/{id} [delete]
func (g gorillaMux) buildDeleteOneCharityMrysAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var act = route.CharityMrysDeleteOne(g.db, g.log, g.ctxTimeout)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("id", vars["id"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// FindBalanceAccounts godoc
// @Summary Find Balance Account By ID
// @Tags Accounts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Account
// @Param account_id path string true "Account ID"
// @Router /v1/accounts/{account_id}/balance [get]
func (g gorillaMux) buildFindBalanceAccountAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewFindBalanceAccountInteractor(
				repository.NewAccountSQL(g.db),
				presenter.NewFindAccountBalancePresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAccountBalanceAction(uc, g.log)
		)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("account_id", vars["account_id"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}

// DeleteAccounts godoc
// @Summary Delete Account By ID
// @Tags Accounts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Account
// @Param id path string true "Account ID"
// @Router /v1/accounts/{id} [delete]
func (g gorillaMux) buildDeleteBalanceAccountAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewDeleteBalanceAccountInteractor(
				repository.NewAccountSQL(g.db),
				presenter.NewDeleteAccountBalancePresenter(),
				g.ctxTimeout,
			)
			act = action.NewDeleteAccountBalanceAction(uc, g.log)
		)

		var (
			vars = mux.Vars(req)
			q    = req.URL.Query()
		)

		q.Add("id", vars["id"])
		req.URL.RawQuery = q.Encode()

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
