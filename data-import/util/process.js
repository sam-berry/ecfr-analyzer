export function failed(err) {
  console.error(`Process failed: ${err}`);
  process.exit(1);
}
