package services

import "tech-challenge-hackaton/internal/infra/clients"

type FFMPEGService struct {
	client *clients.FFMPEGClient
}

func NewFFMPEGService(client *clients.FFMPEGClient) *FFMPEGService {
	return &FFMPEGService{
		client: client,
	}
}

func (f *FFMPEGService) Snapshot(videoID, filename string, interval int) (string, error) {
	// fmt.Println("Processo iniciado:")
	// videoPath := "Marvel_DOTNET_CSHARP.mp4"
	// durationBytes, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", videoPath).Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// durationSeconds, err := strconv.ParseFloat(strings.TrimSpace(string(durationBytes)), 32);
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Duration in seconds %f\n", durationSeconds)
	//
	// interval := 20
	// for curr := 0; curr < int(durationSeconds); curr += interval {
	// 	fmt.Printf("Processando frame: %d\n", curr)
	// 	t := time.Unix(int64(curr), 0).UTC()
	// 	timeFormat := t.Format(time.TimeOnly)
	// 	frameName := fmt.Sprintf("%sframe_at_%s.jpg", videoPath, timeFormat)
	// 	fmt.Println(frameName)
	// 	_, err := exec.Command("ffmpeg", "-ss", timeFormat, "-i", videoPath, "-frames:v", "1", "-q:v", "2", frameName).Output()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	return "", nil
}

