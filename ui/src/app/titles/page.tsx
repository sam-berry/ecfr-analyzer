import { fetchAllTitles } from "ecfr-analyzer/service/EcfrService";
import Error from "ecfr-analyzer/components/Error";

export default async function Page() {
  const titlesResponse = await fetchAllTitles();

  if (titlesResponse.err) {
    return <Error message="An error occurred fetching titles" />;
  }

  return (
    <div>
      {titlesResponse?.data?.titles?.map((title, i) => (
        <div key={i}>{title.name}</div>
      ))}
    </div>
  );
}
