import { auth } from "../gen/v1/auth/auth_pb";
import { Messenger } from "./request";

export namespace AuthService {
  const url = "http://localhost:18080/v1/auth";

  export async function Login(
    data: auth.v1.IAuthRequest
  ): Promise<auth.v1.ILoginResponse> {
    return Messenger.SendRequest<auth.v1.IAuthRequest, auth.v1.ILoginResponse>({
      method: "POST",
      path: "/auth/login",
      data: data,
    });
  }

  export async function Register(
    data: auth.v1.IAuthRequest
  ): Promise<auth.v1.IRegisterResponse> {
    return Messenger.SendRequest<
      auth.v1.IAuthRequest,
      auth.v1.IRegisterResponse
    >({
      method: "POST",
      path: "/auth/register",
      data: data,
    });
  }
}
