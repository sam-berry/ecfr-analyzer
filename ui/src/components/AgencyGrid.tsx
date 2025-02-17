"use client";

import NumberCounter from "ecfr-analyzer/components/NumberCounter";
import InfoPopover from "ecfr-analyzer/components/InfoPopover";
import { ActionIcon, Button, TextInput } from "@mantine/core";
import { IconArrowRight, IconInfoCircle } from "@tabler/icons-react";
import GovInfoBulkDataLink from "ecfr-analyzer/components/GovInfoBulkDataLink";
import { AgencyMetrics } from "ecfr-analyzer/data/AgencyMetrics";
import { useEffect, useState } from "react";
import Fuse from "fuse.js";
import SortButton from "ecfr-analyzer/components/SortButton";
import Link from "next/link";

enum AgencyFilter {
  WORDS_DESC,
  WORDS_ASC,
  SECTIONS_ASC,
  SECTIONS_DESC,
  ALPHA_ASC,
  ALPHA_DESC,
}

const defaultSort = AgencyFilter.WORDS_DESC;

export default function AgencyGrid({
  agencyMetrics,
  disableDetails,
}: {
  agencyMetrics: AgencyMetrics[];
  disableDetails?: boolean;
}) {
  const pageSize = 10;

  const [filter, setFilter] = useState<AgencyFilter>(defaultSort);
  const [filteredAgencies, setFilteredAgencies] = useState<AgencyMetrics[]>([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [visibleCount, setVisibleCount] = useState(pageSize);

  useEffect(() => {
    const fuse = new Fuse(agencyMetrics, {
      keys: [
        "agency.name",
        "agency.shortName",
        "agency.displayName",
        "agency.sortableName",
        "agency.children.name",
      ],
      threshold: 0.3,
    });

    const agencies = (
      searchQuery
        ? fuse.search(searchQuery).map((it) => it.item)
        : agencyMetrics
    ).sort((a, b) => {
      switch (filter) {
        case AgencyFilter.WORDS_DESC:
          return b.metrics.wordCount - a.metrics.wordCount;
        case AgencyFilter.WORDS_ASC:
          return a.metrics.wordCount - b.metrics.wordCount;
        case AgencyFilter.SECTIONS_DESC:
          return b.metrics.sectionCount - a.metrics.sectionCount;
        case AgencyFilter.SECTIONS_ASC:
          return a.metrics.sectionCount - b.metrics.sectionCount;
        case AgencyFilter.ALPHA_ASC:
          return a.agency.name.localeCompare(b.agency.name);
        case AgencyFilter.ALPHA_DESC:
          return b.agency.name.localeCompare(a.agency.name);
        default:
          return 0;
      }
    });

    setFilteredAgencies([...agencies]);
  }, [filter, agencyMetrics, searchQuery]);

  const loadMoreAgencies = () => {
    setVisibleCount((prev) => prev + pageSize);
  };

  const displayedAgencies = filteredAgencies.slice(0, visibleCount);

  return (
    <div className="">
      <div className="m-auto mb-6 w-full max-w-md justify-center">
        <TextInput
          placeholder="Search Agencies"
          value={searchQuery}
          onChange={(event) => setSearchQuery(event.currentTarget.value)}
          classNames={{
            input: "border-primary",
          }}
        />
      </div>
      <div className="flex flex-wrap justify-center gap-4">
        <SortButton
          isAsc={filter === AgencyFilter.WORDS_ASC}
          isDesc={filter === AgencyFilter.WORDS_DESC}
          label="Sort by words"
          sortAsc={() => setFilter(AgencyFilter.WORDS_ASC)}
          sortDesc={() => setFilter(AgencyFilter.WORDS_DESC)}
          clear={() => setFilter(defaultSort)}
        />
        <SortButton
          isAsc={filter === AgencyFilter.SECTIONS_ASC}
          isDesc={filter === AgencyFilter.SECTIONS_DESC}
          label="Sort by regulations"
          sortAsc={() => setFilter(AgencyFilter.SECTIONS_ASC)}
          sortDesc={() => setFilter(AgencyFilter.SECTIONS_DESC)}
          clear={() => setFilter(defaultSort)}
        />
        <SortButton
          isAsc={filter === AgencyFilter.ALPHA_ASC}
          isDesc={filter === AgencyFilter.ALPHA_DESC}
          label="Sort by agency"
          sortAsc={() => setFilter(AgencyFilter.ALPHA_ASC)}
          sortDesc={() => setFilter(AgencyFilter.ALPHA_DESC)}
          clear={() => setFilter(defaultSort)}
        />
      </div>
      <div className="mt-8 flex w-full flex-wrap items-center justify-center gap-12">
        {displayedAgencies.map((it, i) => (
          <Link
            key={i}
            className={`border-primary bg-light w-full max-w-[26rem] shrink-0 border p-4 transition ease-in-out ${disableDetails ? "cursor-default" : "hover:shadow-lg"}`}
            href={disableDetails ? "" : `/agency/${it.agency.slug}`}
            onClick={(event) => {
              if (disableDetails) {
                event.preventDefault();
              }
            }}
          >
            <div
              className="mb-2 line-clamp-2 h-[3.5rem] text-lg font-semibold"
              title={it.agency.name}
            >
              {it.agency.name}
            </div>
            <div className="flex justify-between gap-4">
              {[
                {
                  count: it.metrics.wordCount,
                  emphasize: true,
                  label: "Words",
                  info: (
                    <div>
                      Title words are calculated by selecting all sections
                      attributed to the agency and splitting the text on
                      whitespace, using data available via{" "}
                      <GovInfoBulkDataLink />
                    </div>
                  ),
                },
                {
                  count: it.metrics.sectionCount,
                  label: "Sections",
                  info: (
                    <div>
                      Section count calculated by counting the number of DIV8
                      instances that occur as children of the agency, using data
                      available via <GovInfoBulkDataLink />
                    </div>
                  ),
                },
              ].map((metric, i) => (
                <div
                  key={i}
                  className={`${metric.emphasize ? "text-accent" : "text-primary-700"}`}
                >
                  <div className="text-3xl font-bold uppercase">
                    <NumberCounter
                      start={0}
                      end={metric.count}
                      abbreviate={true}
                    ></NumberCounter>
                  </div>
                  <div className="flex items-center gap-0.5">
                    <div className="font-semibold">{metric.label}</div>
                    <div className="">
                      <InfoPopover
                        target={
                          <ActionIcon size="xs" variant="subtle">
                            <IconInfoCircle
                              size={13}
                              className={`${metric.emphasize ? "text-accent" : "text-primary-700"}`}
                            />
                          </ActionIcon>
                        }
                        width={300}
                      >
                        {metric.info}
                      </InfoPopover>
                    </div>
                  </div>
                </div>
              ))}
            </div>
            {!disableDetails && (
              <div className="-mr-2 mt-4 flex justify-end">
                <Button variant="subtle" size="compact-sm">
                  View details <IconArrowRight size={15} className="ml-1" />
                </Button>
              </div>
            )}
          </Link>
        ))}
      </div>
      {agencyMetrics.length && visibleCount <= agencyMetrics.length && (
        <div className="mt-10 flex justify-center">
          <Button variant="outline" onClick={loadMoreAgencies}>
            Load more agencies
          </Button>
        </div>
      )}
    </div>
  );
}
