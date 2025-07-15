import "@typespec/http";

import "../modules/auth/docs";

using Http;

@service(#{ title: "FirmaERP Server" })
@server(
  "http://127.0.0.1",
  "Um ERP brasileiro de código aberto desenvolvido para pequenas e médias empresas, focado em oferecer funcionalidades essenciais e uma UI simples e intuitiva."
)
namespace FirmaERP;
