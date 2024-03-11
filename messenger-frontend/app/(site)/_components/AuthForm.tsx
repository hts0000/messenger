"use client";

import React, { useCallback, useState } from "react";

import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import { BsGithub, BsGoogle } from "react-icons/bs";

import Input from "@/app/_components/inputs/Input";
import Button from "@/app/_components/Button";
import AuthSocialButton from "./AuthSocialButton";
import { useRouter } from "next/navigation";
import { AuthService } from "@/api/service/auth";
import { toast } from "react-hot-toast";

type AuthFormProps = {};
type Variant = "LOGIN" | "REGISTER";

const AuthForm: React.FC<AuthFormProps> = (props: AuthFormProps) => {
  const router = useRouter();
  const [variant, setVariant] = useState<Variant>("LOGIN");
  const [isLoading, setIsLoading] = useState(false);

  const toggleVariant = useCallback(() => {
    if (variant === "LOGIN") {
      setVariant("REGISTER");
    } else {
      setVariant("LOGIN");
    }
  }, [variant]);

  // react-hook-form
  // defaultValues: 设置form表单中需要提交的字段
  // register: 用于注册表单中需要提交的字段
  // handleSubmit: 接收一个submit函数，给这个处理函数传递从表单中获取的数据
  // formState: 表单状态，errors可以通过register时注册的id拿到那个组件错误
  // 从而进行异常处理
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FieldValues>({
    defaultValues: {
      name: "",
      email: "",
      password: "",
    },
  });

  const onSubmit: SubmitHandler<FieldValues> = async (data) => {
    setIsLoading(true);

    if (variant === "REGISTER") {
      // Axios Register
      AuthService.Register({ ...data })
        .then((resp) => {
          toast.success("register success");
          setVariant("LOGIN");
        })
        .finally(() => {
          setIsLoading(false);
        });
      // axios
      //   .post("http://localhost:18080/v1/auth/register", data)
      //   .then((resp) => {
      //     console.log(resp);
      //     toast.success("register success");
      //     setVariant("LOGIN");
      //   })
      //   .catch((error) => {
      //     toast.error(error.response.data.message);
      //   })
      //   .finally(() => {
      //     setIsLoading(false);
      //   });
    }

    if (variant === "LOGIN") {
      // Axios Login
      AuthService.Login({ ...data })
        .then((resp) => {
          toast.success("login success");
          router.push("/conversations");
        })
        .finally(() => {
          setIsLoading(false);
        });
      // axios
      //   .post("http://localhost:18080/v1/auth/login", data)
      //   .then((resp) => {
      //     console.log(resp);
      //     router.push("/conversations");
      //   })
      //   .catch((error) => {
      //     toast.error(error.response.data.message);
      //   })
      //   .finally(() => {
      //     setIsLoading(false);
      //   });
    }
  };

  const socialAction = (action: string) => {
    setIsLoading(true);
  };

  return (
    <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div className="bg-white px-4 py-8 shadow sm:rounded-lg sm:px-10">
        <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
          {variant === "REGISTER" && (
            <Input
              id="name"
              label="Name"
              register={register}
              errors={errors}
              disabled={isLoading}
            />
          )}
          <Input
            id="email"
            label="Email address"
            register={register}
            errors={errors}
            disabled={isLoading}
          />
          <Input
            id="password"
            label="Password"
            register={register}
            errors={errors}
            disabled={isLoading}
          />
          <div>
            <Button disabled={isLoading} fullWidth type="submit">
              {variant === "LOGIN" ? "Sign in" : "Register"}
            </Button>
          </div>
        </form>
        <div className="mt-6">
          <div className="relative">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-gray-300" />
            </div>
            <div className="relative flex justify-center text-sm">
              <span className="bg-white px-2 text-gray-500">
                Or continue with
              </span>
            </div>
          </div>
        </div>
        <div className="mt-6 flex gap-2">
          <AuthSocialButton
            icon={BsGithub}
            onClick={() => socialAction("github")}
          />
        </div>
        <div className="mt-3 flex gap-2">
          <AuthSocialButton
            icon={BsGoogle}
            onClick={() => socialAction("google")}
          />
        </div>
        <div className="flex gap-2 justify-center text-sm mt-6 px-2 text-gray-500">
          <div>
            {variant === "LOGIN"
              ? "New to Messenger?"
              : "Already have an account?"}
          </div>
          <div onClick={toggleVariant} className="underline cursor-pointer">
            {variant === "LOGIN" ? "Create an account" : "Login"}
          </div>
        </div>
      </div>
    </div>
  );
};

export default AuthForm;
