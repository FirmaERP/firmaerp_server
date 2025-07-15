import "@typespec/http";

using Http;

namespace FirmaERP.Auth;

@tag("Auth")
@route("/auth")
interface Auth {}
