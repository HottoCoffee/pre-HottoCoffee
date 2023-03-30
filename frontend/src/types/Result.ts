export type Result<T, E> = Ok<T> | Err<E>;

interface Ok<T> {
  data: T;
  type: "ok";
}

interface Err<T> {
  data: T;
  type: "err";
}
