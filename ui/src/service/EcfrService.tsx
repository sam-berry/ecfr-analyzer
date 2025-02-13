import {
  errorResponse,
  ResponseContainer,
  successResponse,
} from "ecfr-analyzer/data/ResponseContainer";

const API_ROOT = "https://www.ecfr.gov/api";

const DEFAULT_REVALIDATE = 60 * 60;

export async function fetchAllTitles(): Promise<ResponseContainer<any>> {
  const res = await fetch(`${API_ROOT}/versioner/v1/titles.json`, {
    next: { revalidate: DEFAULT_REVALIDATE },
  });

  if (!res.ok) {
    return errorResponse(`Fetch all titles failed with status: ${res.status}`);
  }

  return successResponse(await res.json());
}
