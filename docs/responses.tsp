import "@typespec/http";

using TypeSpec.Http;

namespace Responses;

model Ok<Body> {
  @body body: Body;
  @statusCode code: 200;
}

model NoContent {
  @statusCode code: 204;
}

model BadRequest {
  @statusCode code: 400;
  message: string;
}

model Unauthorized {
  @statusCode code: 401;
  message: string;
}

model Forbidden {
  @statusCode code: 403;
  message: string;
}

model NotFound {
  @statusCode code: 404;
  message: string;
}

model Internal {
  @statusCode code: 500;
  message: string;
}
