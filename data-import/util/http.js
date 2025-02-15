export function responseContainer({ data, err }) {
  return { data, err };
}

export function successResponse(data) {
  return responseContainer({ data });
}

export function errorResponse({ code, message }) {
  return responseContainer({ err: { code, message } });
}

export function fetchJSON({ url, errorMessage }) {
  return fetch(url, {
    headers: {
      Accept: "application/json",
    },
  })
    .then(async (r) => {
      if (!r.ok) {
        return errorResponse({
          code: r.status,
          message: errorMessage,
        });
      }
      return successResponse(await r.json());
    })
    .catch((err) =>
      errorResponse({ message: `${errorMessage}. Unexpected error: ${err}` }),
    );
}
