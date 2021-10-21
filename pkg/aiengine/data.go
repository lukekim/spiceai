package aiengine

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/spiceai/spiceai/pkg/dataspace"
	"github.com/spiceai/spiceai/pkg/observations"
	"github.com/spiceai/spiceai/pkg/pods"
	"github.com/spiceai/spiceai/pkg/proto/aiengine_pb"
	"github.com/spiceai/spiceai/pkg/state"
)

func SendData(pod *pods.Pod, podState ...*state.State) error {
	if len(podState) == 0 {
		// Nothing to do
		return nil
	}

	err := IsAIEngineHealthy()
	if err != nil {
		return err
	}

	for _, s := range podState {
		addDataRequest := getAddDataRequest(pod, s)

		if addDataRequest == nil {
			continue
		}

		zaplog.Sugar().Debug(aurora.BrightMagenta(fmt.Sprintf("Sending data %d", len(addDataRequest.CsvData))))

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		response, err := aiengineClient.AddData(ctx, addDataRequest)
		if err != nil {
			return fmt.Errorf("failed to post new data to pod %s: %w", pod.Name, err)
		}

		if response.Error {
			return fmt.Errorf("failed to post new data to pod %s: %s", pod.Name, response.Message)
		}

		s.Sent()
	}

	return err
}

func getAddDataRequest(pod *pods.Pod, s *state.State) *aiengine_pb.AddDataRequest {
	if s == nil || !s.TimeSentToAIEngine.IsZero() {
		// Already sent
		return nil
	}

	tagPathMap := pod.TagPathMap()
	categoryPathMap := pod.CategoryPathMap()
	categories := pod.GetDataSpace(s.Path()).Categories()

	csv := strings.Builder{}
	writeHeaders(csv, pod.MeasurementNames(), )

	observations := s.Observations()

	if len(observations) == 0 {
		return nil
	}

	csvPreview := writeData(&csv, pod.Epoch(), s.MeasurementsNames(), tagPathMap[s.Path()], categoryPathMap[s.Path()], observations, 5)

	zaplog.Sugar().Debugf("Posting data to AI engine:\n%s", aurora.BrightYellow(fmt.Sprintf("%s%s...\n%d observations posted", csv.String(), csvPreview, len(observations))))

	addDataRequest := &aiengine_pb.AddDataRequest{
		Pod:     pod.Name,
		CsvData: csv.String(),
	}

	return addDataRequest
}

func writeHeaders(csv *strings.Builder, fqMeasurementsNames []string, categories map[string]*dataspace.Category, fqTags []string) {
	csv.WriteString("time")
	for _, field := range fqMeasurementsNames {
		csv.WriteString(",")
		csv.WriteString(strings.ReplaceAll(field, ".", "_"))
	}
	for fqCategoryName, category := range categories {
		for _, val := range category.Values {
			csv.WriteString(",")
			oneHotFieldName := fmt.Sprintf("%s-%s", fqCategoryName, val)
			oneHotFieldName = strings.ReplaceAll(oneHotFieldName, ".", "_")
			csv.WriteString(oneHotFieldName)
		}
	}
	for _, fqTagName := range fqTags {
		csv.WriteString(",")
		csv.WriteString(strings.ReplaceAll(fqTagName, ".", "_"))
	}
	csv.WriteString("\n")
}

func writeData(csv *strings.Builder, epoch time.Time, measurementsNames []string, categoriesNames []string, tags []string, observations []observations.Observation, previewLines int) string {
	epochTime := epoch.Unix()
	var csvPreview string
	for i, o := range observations {
		if o.Time < epochTime {
			continue
		}

		csv.WriteString(strconv.FormatInt(o.Time, 10))

		// Write measurements		
		for _, f := range measurementsNames {
			csv.WriteString(",")
			measurement, ok := o.Measurements[f]
			if ok {
				csv.WriteString(strconv.FormatFloat(measurement, 'f', -1, 64))
			}
		}

		foundCategories := make(map[string]string)
		for _, category := range categoriesNames {
			if foundVal, ok := foundCategories[category.Name]; ok {
				if foundVal == val {
					foundValMatches = true
				}
			}

			if foundValMatches {
				csv.WriteString("1")
			} else {
				csv.WriteString("0")
			}
		}

		for _, t := range tags {
			csv.WriteString(",")

			hasTag := false
			for _, observationTag := range o.Tags {
				if observationTag == t {
					hasTag = true
					break
				}
			}

			if hasTag {
				csv.WriteString("1")
			} else {
				csv.WriteString("0")
			}
		}
		csv.WriteString("\n")
		if previewLines > 0 && (i+1 == previewLines || (previewLines >= i && i+1 == len(observations))) {
			csvPreview = csv.String()
		}
	}
	return csvPreview
}
