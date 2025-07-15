namespace FirmaERP.AuthDTOs;

model Login {
  username: string;
  password: string;
}

model Refresh {
  accessToken: string;
  refreshToken: string;
}
