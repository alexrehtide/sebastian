import axios from "axios";

type Api = AuthAuthorize &
  AuthAuthenticate &
  RemoteAccountAuthCodeURL &
  RemoteAccountAuthenticate &
  PasswordResettingBegin &
  PasswordResettingEnd &
  RegistrationBegin &
  RegistrationEnd &
  TOTPGenerate;

type AuthAuthorize = (path: "/api/auth/authorize") => Promise<AuthAuthorizeOutput>;
type AuthAuthorizeOutput = { id: number; email: string };

type AuthAuthenticate = (
  path: "/api/auth/authenticate",
  body: AuthAuthenticateInput
) => Promise<AuthAuthenticateOutput>;
type AuthAuthenticateInput = { email: string; password: string };
type AuthAuthenticateOutput = { accessToken: string; refreshToken: string };

type PasswordResettingBegin = (
  path: "/api/password_resetting/begin",
  body: PasswordResettingBeginInput
) => Promise<void>;
type PasswordResettingBeginInput = { email: string };

type PasswordResettingEnd = (path: "/api/password_resetting/end", body: PasswordResettingEndInput) => Promise<void>;
type PasswordResettingEndInput = { resettingCode: string, newPassword: string };

type RegistrationBegin = (path: "/api/registration/begin", body: RegistrationBeginInput) => Promise<void>;
type RegistrationBeginInput = { email: string; password: string; username: string };

type RegistrationEnd = (path: "/api/registration/end", body: RegistrationEndInput) => Promise<RegistrationEndOutput>;
type RegistrationEndInput = { verificationCode: string };
type RegistrationEndOutput = { accessToken: string; refreshToken: string };

type RemoteAccountAuthCodeURL = (
  path: "/api/remote_account/auth_code_url",
  body: RemoteAccountAuthCodeURLInput
) => Promise<RemoteAccountAuthCodeURLOutput>;
type RemoteAccountAuthCodeURLInput = { platform: string };
type RemoteAccountAuthCodeURLOutput = { url: string };

type RemoteAccountAuthenticate = (
  path: "/api/remote_account/authenticate",
  body: RemoteAccountAuthenticateInput
) => Promise<RemoteAccountAuthenticateOutput>;
type RemoteAccountAuthenticateInput = { platform: string; code: string };
type RemoteAccountAuthenticateOutput = { accessToken: string; refreshToken: string };

type TOTPGenerate = (path: "/api/totp/generate") => Promise<TOTPGenerateOutput>;
type TOTPGenerateOutput = { url: string };

const api: Api = (path: string, body: any = {}): any => {
  return axios
    .post(path, body, { headers: { Authorization: localStorage.getItem("ACCESS_TOKEN") ?? undefined } })
    .then((res) => res.data);
};

export default api;
