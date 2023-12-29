import axios from "axios";

type Api = Authorize & Authenticate;

type Authorize = (path: "/api/auth/authorize") => Promise<AuthorizeOutput>;
type AuthorizeOutput = { id: number; email: string };

type Authenticate = (path: "/api/auth/authenticate", body: AuthenticateInput) => Promise<AuthenticateOutput>;
type AuthenticateInput = { email: string; password: string };
type AuthenticateOutput = { accessToken: string; refreshToken: string };

const api: Api = (path: string, body: any = {}): any => {
  return axios
    .post(path, body, { headers: { Authorization: localStorage.getItem("ACCESS_TOKEN") ?? undefined } })
    .then((res) => res.data);
};

export default api;
