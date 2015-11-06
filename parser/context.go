package parser

import "github.com/Xuyuanp/goisp/lexer"

// Context struct
type Context struct {
	parent *Context
	name   string
	token  *lexer.Token
}

// WithContext func
func WithContext(ctx *Context, name string, token *lexer.Token) *Context {
	return &Context{
		parent: ctx,
		name:   name,
		token:  token,
	}
}

// FindToken finds token in context.
func FindToken(ctx *Context, name string) *Token {
	if ctx == nil {
		return nil
	}
	if ctx.name == name {
		return ctx.token
	}
	return FindToken(ctx.parent, name)
}
