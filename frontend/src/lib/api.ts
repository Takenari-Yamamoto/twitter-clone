import { Configuration, AuthApi } from "../api";

const config = new Configuration({
  basePath: "http://localhost:8080/api/v1",
});

export const authApi = new AuthApi(config);
