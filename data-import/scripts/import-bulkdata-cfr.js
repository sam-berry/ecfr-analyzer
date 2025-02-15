import { fetchJSON } from "../util/http.js";
import { failed } from "../util/process.js";
import {
  fetchCFRYearFiles,
  fetchCFRYearTitleFiles,
} from "../service/cfr-service.js";

const importYears = ["2024"];

const yearFiles = await fetchCFRYearFiles(CFR_API_ROOT_URL, importYears);

if (yearFiles.err) {
  failed(yearFiles.err);
}

const example = {
  year: "2024",
  titleArchives: [""],
};

yearFiles.data.reduce((yearFile) => {
  const year = yearFile.name;

  const titleFile = fetchCFRYearTitleFiles(yearFile.url, year);

  return {
    year: year,
    titleArchives: archives,
  };
});

const yearsToTitles = await yearsToImport.reduce(async (acc, year) => {
  const yearResponse = await fetchJSON({
    url: year.link,
    errorMessage: `Failed to fetch bulk data list for ${year}`,
  });

  if (yearResponse.err) {
    failed(yearResponse.err);
  }

  acc[year.name] = yearResponse.data.files.filter((f) => f.folder);
  return acc;
}, {});

Object.keys(yearsToTitles).forEach(async (year) => {
  console.log(`Processing ${year}`);

  const titles = yearsToTitles[year];

  titles.forEach(async (title) => {
    const titleResponse = await fetchJSON({
      url: title.link,
      errorMessage: `Failed to fetch title ${title.name}`,
    });

    if (titleResponse.err) {
      failed(titleResponse.err);
    }

    const titleArchive = titleResponse.data.files.find(
      (f) => f.fileExtension === "zip",
    )?.link;
    console.log(titleArchive);
  });
});

// cfr is historical. ecfr has a single response for each title
