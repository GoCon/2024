export function makePath(path?: string | null): string {
  if (!path) {
    return import.meta.env.BASE_URL;
  }
  return `${import.meta.env.BASE_URL}/${path}`;
}
