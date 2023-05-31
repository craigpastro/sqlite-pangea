import { Database } from "https://deno.land/x/sqlite3@0.9.1/mod.ts";
import {
  assertEquals,
  assertExists,
} from "https://deno.land/std@0.185.0/testing/asserts.ts";

Deno.test("extension", async (t) => {
  const domain = Deno.env.get("PANGEA_DOMAIN");
  assertExists(domain);

  const token = Deno.env.get("PANGEA_TOKEN");
  assertExists(token);

  const db = new Database(":memory:", { enableLoadExtension: true });
  db.loadExtension("./pangea.so");
  db.exec(
    "insert into pangea_config (domain, token) values (?, ?)",
    domain,
    token,
  );

  await t.step("version", async () => {
    const expected = "v" + await Deno.readTextFile("./VERSION");

    const stmt = db.prepare("select pangea_version()");
    const [got] = stmt.value()!;

    assertEquals(got, expected);
  });

  await t.step("redac", () => {
    const stmt = db.prepare(
      `select redact('my phone number is 123-456-7890')`,
    );
    const [got] = stmt.value()!;

    assertEquals(got, "my phone number is <PHONE_NUMBER>");
  });

  await t.step("url reputation", () => {
    const stmt = db.prepare(
      `select url_reputation('https://google.com')`,
    );
    const [got] = stmt.value()!;

    const data = JSON.parse(got as string);
    assertEquals(data.score, -1); // Google's got a good reputation
  });

  await t.step("ip intel", () => {
    const stmt = db.prepare(
      `select ip_intel('23.129.64.211')`,
    );
    const [got] = stmt.value()!;

    const data = JSON.parse(got as string);
    assertEquals(data.reputationData.score, 100); // known suspicious ip
  });

  db.close();
});
