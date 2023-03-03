import { rest, ResponseResolver, RestContext, RestRequest } from "msw";
type A1<T> = T extends (a1: infer I) => unknown ? I : never;

type Method = {
  reqHeaders: any;
  query: any;
  status: number;
  resBody: any;
  reqBody: any;
};
type MethodNames = "get" | "post" | "put" | "patch" | "delete";
type Methods = { [K in MethodNames]: Method };

type Post = {
  post: (option: {
    body: Methods["post"]["reqBody"];
    query: Methods["post"]["query"];
    config?: any;
  }) => Promise<Methods["post"]["resBody"]>;
  $path: () => string;
};

export function restPost<T extends Post>(
  api: T,
  resolver: ResponseResolver<
    RestRequest<A1<T["post"]>["body"]>,
    RestContext,
    Awaited<ReturnType<T["post"]>>["body"]
  >,
) {
  return rest.post(api.$path(), resolver);
}

type Get = {
  get: (option: {
    body: Methods["get"]["reqBody"];
    query: Methods["get"]["query"];
    config?: any;
  }) => Promise<Methods["get"]["resBody"]>;
  $path: () => string;
};

export function restGet<T extends Get>(
  api: T,
  resolver: ResponseResolver<RestRequest<null>, RestContext, Awaited<ReturnType<T["get"]>>["body"]>,
) {
  return rest.get(api.$path(), resolver);
}
