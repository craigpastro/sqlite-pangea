import { Database } from "https://deno.land/x/sqlite3@0.9.1/mod.ts";
import {
  assertEquals,
  assertExists,
} from "https://deno.land/std@0.185.0/testing/asserts.ts";

Deno.test(function testRedac() {
  const token = Deno.env.get("PANGEA_TOKEN");
  assertExists(token);

  const db = new Database(":memory:", { enableLoadExtension: true });

  db.loadExtension("./pangea.so");

  const stmt = db.prepare(
    `select redact('${token}', 'my phone number is 123-456-7890')`,
  );
  const [got] = stmt.value()!;

  assertEquals(got, "my phone number is <PHONE_NUMBER>");

  db.close();
});
