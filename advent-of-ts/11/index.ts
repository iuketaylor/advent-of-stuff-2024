export type Excuse<T> = new (...args: T[]) => {
  [K in keyof T]: `${K & string}: ${T[K] & string}`;
}[keyof T];
