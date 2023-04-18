export const sleep = (millSecond: number) => {
  return new Promise((resolve) => setTimeout(resolve, millSecond));
};
