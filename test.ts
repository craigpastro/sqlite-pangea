import { Database } from "https://deno.land/x/sqlite3@0.9.1/mod.ts";
import {
  assertEquals,
  assertExists,
} from "https://deno.land/std@0.185.0/testing/asserts.ts";

Deno.test(function testRedac() {
  const token = Deno.env.get("PANGEA_TOKEN");
  assertExists(token);

  const db = new Database("redac", { enableLoadExtension: true, memory: true });
  db.loadExtension("./pangea.so");

  const stmt = db.prepare(
    `select redact('${token}', 'my phone number is 123-456-7890')`,
  );
  const [got] = stmt.value()!;

  db.close();

  assertEquals(got, "my phone number is <PHONE_NUMBER>");
});

Deno.test(function testUrlReputation() {
  const token = Deno.env.get("PANGEA_TOKEN");
  assertExists(token);

  const db = new Database("url_reputation", {
    enableLoadExtension: true,
    memory: true,
  });
  db.loadExtension("./pangea.so");

  const stmt = db.prepare(
    `select url_reputation('${token}', 'https://google.com')`,
  );
  const [got] = stmt.value()!;
  db.close();

  const data = JSON.parse(got as string);
  assertEquals(data.score, -1); // Google's got a good reputation
});
