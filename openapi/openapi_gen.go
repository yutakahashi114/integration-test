// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// ID defines model for ID.
type ID struct {
	Id int64 `json:"id"`
}

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody User

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody User

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 全ユーザー取得
	// (GET /user)
	FindUsers(w http.ResponseWriter, r *http.Request)
	// ユーザー作成
	// (POST /user)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// 指定したIDのユーザー削除
	// (DELETE /user/{id})
	DeleteUser(w http.ResponseWriter, r *http.Request, id int64)
	// 指定したIDのユーザー取得
	// (GET /user/{id})
	GetUserByID(w http.ResponseWriter, r *http.Request, id int64)
	// 指定したIDのユーザー更新
	// (PUT /user/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// FindUsers operation middleware
func (siw *ServerInterfaceWrapper) FindUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindUsers(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUserByID operation middleware
func (siw *ServerInterfaceWrapper) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserByID(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user", wrapper.FindUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/user", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/user/{id}", wrapper.DeleteUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user/{id}", wrapper.GetUserByID)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/user/{id}", wrapper.UpdateUser)
	})

	return r
}
