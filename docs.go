// Copyright 2022 Slawomir Caluch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package ctxslog provides access to go's
log/slog logging from a [context.Context].

It defines a set of functions to hadle loggers
in a context:
- [NewContext] to create a new context with a logger
- [FromContext] to retrieve the logger from a context
*/
package ctxslog
