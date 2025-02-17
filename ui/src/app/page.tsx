import { fetchTitleMetrics } from "ecfr-analyzer/service/MetricService";
import Error from "ecfr-analyzer/components/Error";
import Navigation from "ecfr-analyzer/components/Navigation";
import InfoPopover from "ecfr-analyzer/components/InfoPopover";
import { ActionIcon } from "@mantine/core";
import { IconInfoCircle } from "@tabler/icons-react";
import NumberCounter from "ecfr-analyzer/components/NumberCounter";

const govInfoLink = (
  <a
    href="https://www.govinfo.gov/bulkdata/ECFR"
    className="font-medium underline"
    target="_blank"
  >
    GovInfo ECFR Bulk Data
  </a>
);

export default async function Page() {
  const titleMetricsResponse = await fetchTitleMetrics();
  const titleMetrics = titleMetricsResponse.data;

  if (titleMetricsResponse.err || !titleMetrics) {
    return <Error message={titleMetricsResponse.err?.message} />;
  }

  return (
    <>
      <Navigation />
      <div className="mt-24 flex justify-center px-4">
        <div className="relative w-full max-w-3xl">
          <div className="bg-primary absolute -left-4 -top-4 bottom-4 right-4"></div>
          <div className="bg-accent absolute -bottom-4 -right-4 left-4 top-4"></div>
          <div className="bg-light border-primary relative flex flex-wrap justify-evenly gap-8 border p-8">
            {[
              {
                count: titleMetrics.wordCount,
                label: "Total Words",
                info: (
                  <div>
                    Word count calculated by extracting the text from each title
                    BODY content and splitting on whitespace, using data
                    available via {govInfoLink}
                  </div>
                ),
              },
              {
                count: titleMetrics.sectionCount,
                label: "Total Sections",
                info: (
                  <div>
                    Section count calculated by counting the number of DIV8
                    instances in each title, using data available via{" "}
                    {govInfoLink}
                  </div>
                ),
              },
            ].map((it, i) => (
              <div
                key={i}
                className="flex shrink-0 flex-col items-center gap-2"
              >
                <div className="font-title text-5xl font-bold">
                  <NumberCounter start={0} end={it.count}></NumberCounter>
                </div>
                <div className="flex items-center gap-0.5">
                  <div className="font-semibold">{it.label}</div>
                  <div className="-mr-4">
                    <InfoPopover
                      target={
                        <ActionIcon size="xs" variant="subtle">
                          <IconInfoCircle size={13} />
                        </ActionIcon>
                      }
                      width={300}
                    >
                      {it.info}
                    </InfoPopover>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
