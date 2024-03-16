import { makePath } from "../../src/utils/makePath";
import { describe, assert, expect, test } from "vitest";

describe("makePath", () => {
  test("prefixにBASE_URLを追加する", () => {
    expect(makePath("path")).eq(`${import.meta.env.BASE_URL}/path`);
  });

  test("空文字が渡されたらBASE_URLだけを返す", () => {
    expect(makePath("")).eq(`${import.meta.env.BASE_URL}`);
  });

  test("NULLが渡されたらBASE_URLだけを返す", () => {
    expect(makePath(null)).eq(`${import.meta.env.BASE_URL}`);
  });

  test("引数がなかったらBASE_URLだけを返す", () => {
    expect(makePath()).eq(`${import.meta.env.BASE_URL}`);
  });
});
