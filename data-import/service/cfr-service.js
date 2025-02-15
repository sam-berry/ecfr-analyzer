import { errorResponse, fetchJSON, successResponse } from "../util/http.js";

export async function fetchCFRYearFiles(url, includeYears) {
  const allYears = await fetchJSON({
    url: url,
    errorMessage: `Failed to fetch export years, URL: ${url}`,
  });

  if (allYears.err) {
    return errorResponse(allYears.err);
  }

  const filteredYears = allYears.data.files.filter((y) =>
    includeYears.includes(y.name),
  );

  return successResponse(filteredYears);
}

export async function fetchCFRYearTitleFiles(url, year) {
  const res = await fetchJSON({
    url: url,
    errorMessage: `Failed to fetch title for ${year}, URL: ${url}`,
  });

  if (res.err) {
    return errorResponse(res.err);
  }

  return successResponse(res.data.files);
}
