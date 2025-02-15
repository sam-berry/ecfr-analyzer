import {
  fetchFullTitle,
  fetchSearchResults,
} from "ecfr-analyzer/service/EcfrService";
import Error from "ecfr-analyzer/components/Error";

function countWords(str: string): number {
  return str.split(" ").filter(function (n) {
    return n != "";
  }).length;
}

const delay = (delayInms) => {
  return new Promise((resolve) => setTimeout(resolve, delayInms));
};

export default async function Page(props: {
  params: Promise<{ slug: string }>;
}) {
  const { slug } = await props.params;

  const sr = await fetchSearchResults(slug);

  if (sr.err) {
    return <Error message={sr.err.message} />;
  }

  console.log(sr.data.results);

  const results = await Promise.all(
    sr.data?.results?.map(async (result, i) => {
      const hierarchy = result.hierarchy;

      if (i > 0) {
        await delay(10000);
      }

      return fetchFullTitle("2025-01-01", hierarchy).then((r) => {
        if (r.err) {
          console.log(r.err);
          return;
        }

        console.log(r.data);

        const titleData = r.data?.replace(
          /\<\?xml.+\?\>|\<\!DOCTYPE.+]\>/g,
          "",
        );

        const wordCount = countWords(titleData);
        console.log(`${hierarchy.title} : ${wordCount}`);

        return wordCount;
      });
    }),
  );

  console.log("Results: ", results);

  return (
    <div>
      <div>Total Words: {results?.reduce((acc, cur) => acc + cur, 0)}</div>
      <div>Total Sections: {results?.length}</div>
    </div>
  );
}
