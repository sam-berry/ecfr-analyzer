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
    return errorResponse({
      code: res.status,
      message: "An error occurred fetching titles",
    });
  }

  return successResponse(await res.json());
}

export async function fetchAllAgencies(): Promise<ResponseContainer<any>> {
  const res = await fetch(`${API_ROOT}/admin/v1/agencies.json`, {
    next: { revalidate: DEFAULT_REVALIDATE },
  });

  if (!res.ok) {
    return errorResponse({
      code: res.status,
      message: "An error occurred fetching agencies",
    });
  }

  return successResponse(await res.json());
}

export async function fetchSearchResults(
  slug: string,
): Promise<ResponseContainer<any>> {
  const res = await fetch(
    `${API_ROOT}/search/v1/results?${new URLSearchParams({
      "agency_slugs[]": slug,
      per_page: "1000",
      page: "1",
      order: "relevance",
      paginate_by: "results",
    }).toString()}`,
    {
      cache: "no-cache",
    },
  );

  if (!res.ok) {
    return errorResponse({
      code: res.status,
      message: "An error occurred fetching search results for agency",
    });
  }

  return successResponse(await res.json());
}

function removeNullProperties(data: any) {
  return Object.fromEntries(
    Object.entries(data).filter(([, value]) => value !== null),
  );
}

export async function fetchFullTitle(
  date: string,
  data: any,
): Promise<ResponseContainer<any>> {
  const title = data.title;

  const res = await fetch(
    `${API_ROOT}/versioner/v1/full/${date}/title-${title}.xml?${new URLSearchParams(
      {
        date,
        ...removeNullProperties(data),
      },
    ).toString()}`,
    {
      cache: "no-cache",
    },
  );

  if (!res.ok) {
    return errorResponse({
      code: res.status,
      message: `An error occurred fetching full title: ${title}`,
    });
  }

  return successResponse(await res.text());
}
