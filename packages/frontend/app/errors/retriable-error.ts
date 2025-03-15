export class RetriableError extends Error {
  public readonly name = "RetriableError";

  constructor() {
    super();
  }
}
