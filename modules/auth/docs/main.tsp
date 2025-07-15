import "@typespec/http";

import "./dtos.tsp";
import "./models.tsp";
import "../../../docs/responses.tsp";

using Http;

namespace FirmaERP.Auth;

@tag("Auth")
@route("/auth")
@doc("Authentication module")
interface Auth {
  @post
  @route("/login")
  login(@body dto: AuthDTOs.Login):
    | Responses.Ok<AuthModels.Token>
    | Responses.BadRequest
    | Responses.Unauthorized
    | Responses.Internal;

  @post
  @route("/refresh")
  refresh(@body dto: AuthDTOs.Refresh):
    | Responses.Ok<AuthModels.Token>
    | Responses.BadRequest
    | Responses.Forbidden
    | Responses.Unauthorized
    | Responses.Internal;
}
