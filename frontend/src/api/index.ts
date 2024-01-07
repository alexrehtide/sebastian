import axios from "axios";

type Api = AuthAuthorize &
  AuthAuthenticate &
  AuthBeginRegistration &
  AuthEndRegistration &
  Oauth2AuthCodeURL &
  Oauth2Authenticate &
  TOTPGenerate;

type AuthAuthorize = (path: "/api/auth/authorize") => Promise<AuthAuthorizeOutput>;
type AuthAuthorizeOutput = { id: number; email: string };

type AuthAuthenticate = (
  path: "/api/auth/authenticate",
  body: AuthAuthenticateInput
) => Promise<AuthAuthenticateOutput>;
type AuthAuthenticateInput = { email: string; password: string };
type AuthAuthenticateOutput = { accessToken: string; refreshToken: string };

type AuthBeginRegistration = (
  path: "/api/auth/begin_registration",
  body: AuthBeginRegistrationInput
) => Promise<void>;
type AuthBeginRegistrationInput = { email: string; password: string; username: string };

type AuthEndRegistration = (
  path: "/api/auth/end_registration",
  body: AuthEndRegistrationInput
) => Promise<AuthEndRegistrationOutput>;
type AuthEndRegistrationInput = { verificationCode: string };
type AuthEndRegistrationOutput = { accessToken: string; refreshToken: string };

type Oauth2AuthCodeURL = (
  path: "/api/oauth2/auth_code_url",
  body: Oauth2AuthCodeURLInput
) => Promise<Oauth2AuthCodeURLOutput>;
type Oauth2AuthCodeURLInput = { platform: string };
type Oauth2AuthCodeURLOutput = { url: string };

type Oauth2Authenticate = (
  path: "/api/oauth2/authenticate",
  body: Oauth2AuthenticateInput
) => Promise<Oauth2AuthenticateOutput>;
type Oauth2AuthenticateInput = { platform: string; code: string };
type Oauth2AuthenticateOutput = { accessToken: string; refreshToken: string };

type TOTPGenerate = (path: "/api/totp/generate") => Promise<TOTPGenerateOutput>;
type TOTPGenerateOutput = { url: string };

const api: Api = (path: string, body: any = {}): any => {
  return axios
    .post(path, body, { headers: { Authorization: localStorage.getItem("ACCESS_TOKEN") ?? undefined } })
    .then((res) => res.data);
};

export default api;
