package translated

import "github.com/codelingo/php/passes/togo/internal/phpctx"

func Echo(ctx phpctx.PHPContext) {
	ctx.Echo.Write("test")
}
