import axios from "axios";

type Api = AuthAuthorize & AuthAuthenticate & TOTPGenerate;

type AuthAuthorize = (path: "/api/auth/authorize") => Promise<AuthAuthorizeOutput>;
type AuthAuthorizeOutput = { id: number; email: string };

type AuthAuthenticate = (path: "/api/auth/authenticate", body: AuthAuthenticateInput) => Promise<AuthAuthenticateOutput>;
type AuthAuthenticateInput = { email: string; password: string };
type AuthAuthenticateOutput = { accessToken: string; refreshToken: string };

type TOTPGenerate = (path: "/api/totp/generate") => Promise<TOTPGenerateOutput>
type TOTPGenerateOutput = { url: string }

const api: Api = (path: string, body: any = {}): any => {
  return axios
    .post(path, body, { headers: { Authorization: localStorage.getItem("ACCESS_TOKEN") ?? undefined } })
    .then((res) => res.data);
};

export default api;
