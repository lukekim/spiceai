import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';

import PodHeader from '../components/app/PodHeader';
import Card from '../components/layout/Card';
import { usePod } from '../models/pod';
import { useFlights } from '../models/flight';
import { useObservations } from '../models/observation';
import FlightChart from '../components/flights/FlightChart';
import DataEditor, {
  DataEditorContainer,
  GridColumn,
  GridCell,
  GridCellKind,
} from '@glideapps/glide-data-grid';

interface PodProps {
  podName: string;
}

// TODO: Resize dynamically
const gridWidth = 900;
const gridHeight = 600;

const PodPage: React.FunctionComponent<PodProps> = () => {
  const location = useLocation();
  const podNamePathIndex = location.pathname.lastIndexOf('/') + 1;
  const podName = location.pathname.substring(podNamePathIndex);

  const { data: pod, error: podError } = usePod(podName);
  const { data: flights, error: flightsError } = useFlights(podName);
  const { data: observations, error: observationsError } = useObservations(podName);

  const [gridColumns, setGridColumns] = useState<GridColumn[]>([]);

  useEffect(() => {
    if (pod) {
      const cols: GridColumn[] = [{ title: 'time', width: 180 }];
      const numMeasurements = pod.measurements ? pod.measurements.length : 0
      const numCategories = pod.categories ? pod.categories.length : 0

      const colWidth =
        (gridWidth - cols[0].width - 17) / (numMeasurements + numCategories + 1);

      if (pod.measurements) {
        for (const m of pod.measurements) {
          cols.push({ title: m, width: colWidth });
        }
      }

      if (pod.categories) {
        for (const c of pod.categories) {
          cols.push({ title: c, width: colWidth });
        }
      }

      cols.push({ title: 'tags', width: colWidth });

      setGridColumns(cols);
    }
  }, [pod]);

  const getGridDataFunc = ([col, row]: readonly [number, number]): GridCell => {
    if (!pod || !observations) {
      return {
        kind: GridCellKind.Number,
        data: row,
        displayData: row.toString(),
        allowOverlay: false,
      };
    }
    
    if (row >= observations.length) {
      return {
        kind: GridCellKind.Number,
        data: undefined,
        displayData: '',
        allowOverlay: false,
      };
    }
    const observation = observations[observations.length - row - 1];
    if (col === 0) {
      return {
        kind: GridCellKind.Number,
        data: observation.time,
        displayData: new Date(observation.time * 1000).toLocaleString(),
        allowOverlay: false,
      };
    }

    console.log(observation)

    const numMeasurements = pod.measurements ? pod.measurements.length : 0
    const numCategories = pod.categories ? pod.categories.length : 0

    if (observation.measurements && pod.measurements && col >= 1 && col <= numMeasurements) {
      const measurement = observation.measurements[col - 1];
      return {
        kind: GridCellKind.Number,
        data: measurement,
        displayData: measurement ? measurement.toString() : "",
        allowOverlay: false,
      };
    }

    if (observation.categories && col > numMeasurements && col <= numMeasurements + numCategories) {
      const category =
        observation.categories[pod.categories[col - numMeasurements - 1]];
      return {
        kind: GridCellKind.Text,
        data: category,
        displayData: category,
        allowOverlay: false,
      };
    }

    if (col == numMeasurements + numCategories + 1) {
      const tags = observation.tags ? observation.tags.join(' ') : '';
      return {
        kind: GridCellKind.Text,
        data: tags,
        displayData: tags,
        allowOverlay: false,
      };
    }

    return {
      kind: GridCellKind.Number,
      data: row,
      displayData: row.toString(),
      allowOverlay: false,
    };
  };

  return (
    <div className="flex flex-col flex-grow">
      {!podError && pod && (
        <div className="mb-2">
          <PodHeader pod={pod}></PodHeader>
          <h2 className="ml-2 mb-2 font-spice tracking-spice text-s uppercase">Observations</h2>
          {observationsError && (
            <span>An error occurred fetching observations: {observationsError}</span>
          )}
          <div className="border-1 border-gray-300">
            {!observationsError && observations && gridColumns && (
              <DataEditorContainer width={gridWidth} height={gridHeight}>
                <DataEditor
                  getCellContent={getGridDataFunc}
                  columns={gridColumns}
                  rows={observations.length}
                  rowMarkers={false}
                />
              </DataEditorContainer>
            )}
          </div>
          <h2 className="mt-4 ml-2 mb-2 font-spice tracking-spice text-s uppercase">
            Training Runs
          </h2>
          <div className="p-2">
            {!flightsError &&
              flights.map((flight, i) => (
                <div key={i}>
                  <Card>
                    <FlightChart flight={flight} />
                  </Card>
                </div>
              ))}
            {(!flights || flights.length === 0) && <span>Pod has no training runs.</span>}
          </div>
        </div>
      )}
    </div>
  );
};

export default PodPage;
