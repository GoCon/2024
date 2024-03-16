// パスを組み立てるときに環境変数を使ってルート情報を追加する関数。
// GitHub Pages で公開する都合で astro.config.ts で base を指定している。
// Netlifyで公開するときはbaseを上書きして公開しているので動的に指定する必要がある。
export function makePath(path?: string | null): string {
  if (!path) {
    return import.meta.env.BASE_URL;
  }
  return `${import.meta.env.BASE_URL}/${path}`;
}
