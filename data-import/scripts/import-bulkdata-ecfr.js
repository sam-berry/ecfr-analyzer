import { fetchJSON } from "../util/http.js";

const rootURL = "https://www.govinfo.gov/bulkdata/json/ECFR";
const importYears = ["2024"];

const years = await fetchJSON({
  url: rootURL,
  errorMessage: "Failed to fetch export years",
});

if (years.err) {
  console.error(years.err);
  process.exit(1);
}

console.log(JSON.stringify(years.data));

const yearsToImport = years.data.files.filter((y) =>
  importYears.includes(y.name),
);

yearsToImport.map(async (year) => {
  const yearResponse = await fetchJSON({
    url: year.link,
    errorMessage: `Failed to fetch bulk data list for ${year}`,
  });

  if (yearResponse.err) {
    console.error(yearResponse.err);
    process.exit(1);
  }

  console.log(JSON.stringify(yearResponse.data));
});
