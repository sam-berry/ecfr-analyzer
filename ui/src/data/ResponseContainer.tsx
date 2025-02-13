export interface ResponseContainer<T> {
  data?: T;
  err?: any;
}

export function successResponse<T>(data: T): ResponseContainer<T> {
  return { data };
}

export function errorResponse<T>(err: any): ResponseContainer<T> {
  return { err };
}
