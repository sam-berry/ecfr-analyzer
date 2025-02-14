import { fetchAllTitles } from "ecfr-analyzer/service/EcfrService";
import Error from "ecfr-analyzer/components/Error";

export default async function Page() {
  const titlesResponse = await fetchAllTitles();

  if (titlesResponse.err) {
    return <Error message={titlesResponse.err.message} />;
  }

  return (
    <div>
      {titlesResponse?.data?.titles?.map((it, i) => (
        <div key={i}>{it.name}</div>
      ))}
    </div>
  );
}
