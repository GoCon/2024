import { makePath } from "../../src/utils/makePath";
import { describe, assert, expect, test } from "vitest";

describe("makePath", () => {
  test("success", () => {
    expect(makePath("path")).eq(`${import.meta.env.BASE_URL}/path`);
  });
});
