import { concatWithBase } from "../../src/utils/concatWithBase";
import { describe, expect, test } from "vitest";

describe("concatWithBase", () => {
  test("prefixにBASE_URLを追加する", () => {
    expect(concatWithBase("path")).eq(`${import.meta.env.BASE_URL}/path`);
  });

  test("空文字が渡されたらBASE_URLだけを返す", () => {
    expect(concatWithBase("")).eq(`${import.meta.env.BASE_URL}`);
  });

  test("引数がなかったらBASE_URLだけを返す", () => {
    expect(concatWithBase()).eq(`${import.meta.env.BASE_URL}`);
  });
});
