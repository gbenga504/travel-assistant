export function omit<T extends object, K extends keyof T>(
  obj: T,
  paths: K[]
): Omit<T, K> {
  return Object.fromEntries(
    Object.entries(obj).filter(([key]) => !paths.includes(key as K))
  ) as Omit<T, K>;
}
