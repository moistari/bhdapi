// Package bhdapi provides a bhd client.
package bhdapi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

// SearchRequest is a bhd search request.
type SearchRequest struct {
	// The torrent name. It does support !negative searching. Example: Christmas Movie
	Search string `json:"search,omitempty"`
	// The torrent info_hash. This is an exact match.
	InfoHash string `json:"info_hash,omitempty"`
	// The torrent folder name. This is an exact match.
	FolderName string `json:"folder_name,omitempty"`
	// The torrent included file names. This is an exact match.
	FileName string `json:"file_name,omitempty"`
	// The torrent size. This is an exact match.
	Size int64 `json:"size,omitempty"`
	// The uploaders username. Only non anonymous results will be returned.
	UploadedBy string `json:"uploaded_by,omitempty"`
	// The ID of the matching IMDB page.
	ImdbID string `json:"imdb_id,omitempty"`
	// The ID of the matching TMDB page.
	TmdbID string `json:"tmdb_id,omitempty"`
	// Any categories separated by comma(s). (TV, Movies)
	Categories []string `json:"categories,omitempty"`
	// Any types separated by comma(s). (BD Remux, 1080p, etc.)
	Types []string `json:"types,omitempty"`
	// Any sources separated by comma(s). (Blu-ray, WEB, DVD, etc.)
	Sources []string `json:"sources,omitempty"`
	// Any genres separated by comma(s). (Action, Anime, Stand-Up, Western, etc.)
	Genres []string `json:"genres,omitempty"`
	// Any internal release groups separated by comma(s). (FraMeSToR, BHDStudio, BeyondHD, RPG, iROBOT, iFT, ZR, MKVULTRA)
	Groups []string `json:"groups,omitempty"`
	// The torrent freeleech status. 1 = Must match.
	Freeleech Bool `json:"freeleech,omitempty"`
	// The torrent limited UL promo. 1 = Must match.
	Limited Bool `json:"limited,omitempty"`
	// The torrent 25% promo. 1 = Must match.
	Promo25 Bool `json:"promo25,omitempty"`
	// The torrent 50% promo. 1 = Must match.
	Promo50 Bool `json:"promo50,omitempty"`
	// The torrent 75% promo. 1 = Must match.
	Promo75 Bool `json:"promo75,omitempty"`
	// The torrent refund promo. 1 = Must match.
	Refund Bool `json:"refund,omitempty"`
	// The torrent rescue promo. 1 = Must match.
	Rescue Bool `json:"rescue,omitempty"`
	// The torrent rewind promo. 1 = Must match.
	Rewind Bool `json:"rewind,omitempty"`
	// The torrent Stream Optimized flag. 1 = Must match.
	Stream Bool `json:"stream,omitempty"`
	// The torrent SD flag. 1 = Must match.
	SD Bool `json:"sd,omitempty"`
	// The torrent TV pack flag. 1 = Must match.
	Pack Bool `json:"pack,omitempty"`
	// The torrent x264/h264 codec flag. 1 = Must match.
	H264 Bool `json:"h_264,omitempty"`
	// The torrent x265/h265 codec flag. 1 = Must match.
	H265 Bool `json:"h_265,omitempty"`
	// Any features separated by comma(s). (DV, HDR10, HDR10P, Commentary)
	Features []string `json:"features,omitempty"`
	// The torrent has at least 1 seeder. 1 = Must match.
	Alive Bool `json:"alive,omitempty"`
	// The torrent has less than 3 seeders. 1 = Must match.
	Dying Bool `json:"dying,omitempty"`
	// The torrent has no seeders. 1 = Must match.
	Dead Bool `json:"dead,omitempty"`
	// The torrent has no seeders and an active reseed request. 1 = Must match.
	Reseed Bool `json:"reseed,omitempty"`
	// The torrent is seeded by you. 1 = Must match.
	Seeding Bool `json:"seeding,omitempty"`
	// The torrent is being leeched by you. 1 = Must match.
	Leeching Bool `json:"leeching,omitempty"`
	// The torrent has been completed by you. 1 = Must match.
	Completed Bool `json:"completed,omitempty"`
	// The torrent has not been completed by you. 1 = Must match.
	Incomplete Bool `json:"incomplete,omitempty"`
	// The torrent has not been downloaded you. 1 = Must match.
	NotDownloaded Bool `json:"notdownloaded,omitempty"`
	// The minimum BHD rating.
	MinBHD int `json:"min_bhd,omitempty"`
	// The minimum number of BHD votes.
	VoteBHD int `json:"vote_bhd,omitempty"`
	// The minimum IMDb rating.
	MinImdb int `json:"min_imdb,omitempty"`
	// The minimum number of IMDb votes.
	VoteImdb int `json:"vote_imdb,omitempty"`
	// The minimum TMDb rating.
	MinTmbd int `json:"min_tmdb,omitempty"`
	// The minimum number of TDMb votes.
	VoteTmbd int `json:"vote_tmdb,omitempty"`
	// The earliest release year.
	MinYear int `json:"min_year,omitempty"`
	// The latest release year.
	MaxYear int `json:"max_year,omitempty"`
	// Any production countries separated by comma(s). (France, Japan, etc.)
	Countries []string `json:"countries,omitempty"`
	// Any spoken languages separated by comma(s). (French, English, etc.)
	Languages []string `json:"languages,omitempty"`
	// Any audio tracks separated by comma(s). (English, Japanese, etc.)
	Audios []string `json:"audios,omitempty"`
	// Any subtitles separated by comma(s). (Dutch, Finnish, Swedish, etc.)
	Subtitles []string `json:"subtitles,omitempty"`
	// Field to sort results by. (bumped_at, created_at, seeders, leechers, times_completed, size, name, imdb_rating, tmdb_rating, bhd_rating). Default is bumped_at
	Sort string `json:"sort,omitempty"`
	// The direction of the sort of results. (asc, desc). Default is desc
	Order string `json:"order,omitempty"`
	// The page number of the results. Only if the result set has more than 100 total matches.
	Page int `json:"page,omitempty"`
}

func Search(query ...string) *SearchRequest {
	return &SearchRequest{
		Search: strings.Join(query, " "),
	}
}

func (req SearchRequest) WithInfoHash(infoHash string) *SearchRequest {
	req.InfoHash = infoHash
	return &req
}

func (req SearchRequest) WithFolderName(folderName string) *SearchRequest {
	req.FolderName = folderName
	return &req
}

func (req SearchRequest) WithFileName(fileName string) *SearchRequest {
	req.FileName = fileName
	return &req
}

func (req SearchRequest) WithSize(size int64) *SearchRequest {
	req.Size = size
	return &req
}

func (req SearchRequest) WithUploadedBy(uploadedBy string) *SearchRequest {
	req.UploadedBy = uploadedBy
	return &req
}

func (req SearchRequest) WithImdbID(imdbID string) *SearchRequest {
	req.ImdbID = imdbID
	return &req
}

func (req SearchRequest) WithTmdbID(tmdbID string) *SearchRequest {
	req.TmdbID = tmdbID
	return &req
}

func (req SearchRequest) WithCategories(categories ...string) *SearchRequest {
	req.Categories = categories
	return &req
}

func (req SearchRequest) WithTypes(types ...string) *SearchRequest {
	req.Types = types
	return &req
}

func (req SearchRequest) WithSources(sources ...string) *SearchRequest {
	req.Sources = sources
	return &req
}

func (req SearchRequest) WithGenres(genres ...string) *SearchRequest {
	req.Genres = genres
	return &req
}

func (req SearchRequest) WithGroups(groups ...string) *SearchRequest {
	req.Groups = groups
	return &req
}

func (req SearchRequest) WithFreeleech(freeleech bool) *SearchRequest {
	req.Freeleech = Bool(freeleech)
	return &req
}

func (req SearchRequest) WithLimited(limited bool) *SearchRequest {
	req.Limited = Bool(limited)
	return &req
}

func (req SearchRequest) WithPromo25(promo25 bool) *SearchRequest {
	req.Promo25 = Bool(promo25)
	return &req
}

func (req SearchRequest) WithPromo50(promo50 bool) *SearchRequest {
	req.Promo50 = Bool(promo50)
	return &req
}

func (req SearchRequest) WithPromo75(promo75 bool) *SearchRequest {
	req.Promo75 = Bool(promo75)
	return &req
}

func (req SearchRequest) WithRefund(refund bool) *SearchRequest {
	req.Refund = Bool(refund)
	return &req
}

func (req SearchRequest) WithRescue(rescue bool) *SearchRequest {
	req.Rescue = Bool(rescue)
	return &req
}

func (req SearchRequest) WithRewind(rewind bool) *SearchRequest {
	req.Rewind = Bool(rewind)
	return &req
}

func (req SearchRequest) WithStream(stream bool) *SearchRequest {
	req.Stream = Bool(stream)
	return &req
}

func (req SearchRequest) WithSD(sd bool) *SearchRequest {
	req.SD = Bool(sd)
	return &req
}

func (req SearchRequest) WithPack(pack bool) *SearchRequest {
	req.Pack = Bool(pack)
	return &req
}

func (req SearchRequest) WithH264(h264 bool) *SearchRequest {
	req.H264 = Bool(h264)
	return &req
}

func (req SearchRequest) WithH265(h265 bool) *SearchRequest {
	req.H265 = Bool(h265)
	return &req
}

func (req SearchRequest) WithFeatures(features ...string) *SearchRequest {
	req.Features = features
	return &req
}

func (req SearchRequest) WithAlive(alive bool) *SearchRequest {
	req.Alive = Bool(alive)
	return &req
}

func (req SearchRequest) WithDying(dying bool) *SearchRequest {
	req.Dying = Bool(dying)
	return &req
}

func (req SearchRequest) WithDead(dead bool) *SearchRequest {
	req.Dead = Bool(dead)
	return &req
}

func (req SearchRequest) WithReseed(reseed bool) *SearchRequest {
	req.Reseed = Bool(reseed)
	return &req
}

func (req SearchRequest) WithSeeding(seeding bool) *SearchRequest {
	req.Seeding = Bool(seeding)
	return &req
}

func (req SearchRequest) WithLeeching(leeching bool) *SearchRequest {
	req.Leeching = Bool(leeching)
	return &req
}

func (req SearchRequest) WithCompleted(completed bool) *SearchRequest {
	req.Completed = Bool(completed)
	return &req
}

func (req SearchRequest) WithIncomplete(incomplete bool) *SearchRequest {
	req.Incomplete = Bool(incomplete)
	return &req
}

func (req SearchRequest) WithNotDownloaded(notDownloaded bool) *SearchRequest {
	req.NotDownloaded = Bool(notDownloaded)
	return &req
}

func (req SearchRequest) WithMinBHD(minBHD int) *SearchRequest {
	req.MinBHD = minBHD
	return &req
}

func (req SearchRequest) WithVoteBHD(voteBHD int) *SearchRequest {
	req.VoteBHD = voteBHD
	return &req
}

func (req SearchRequest) WithMinImdb(minImdb int) *SearchRequest {
	req.MinImdb = minImdb
	return &req
}

func (req SearchRequest) WithVoteImdb(voteImdb int) *SearchRequest {
	req.VoteImdb = voteImdb
	return &req
}

func (req SearchRequest) WithMinTmbd(minTmbd int) *SearchRequest {
	req.MinTmbd = minTmbd
	return &req
}

func (req SearchRequest) WithVoteTmbd(voteTmbd int) *SearchRequest {
	req.VoteTmbd = voteTmbd
	return &req
}

func (req SearchRequest) WithMinYear(minYear int) *SearchRequest {
	req.MinYear = minYear
	return &req
}

func (req SearchRequest) WithMaxYear(maxYear int) *SearchRequest {
	req.MaxYear = maxYear
	return &req
}

func (req SearchRequest) WithCountries(countries ...string) *SearchRequest {
	req.Countries = countries
	return &req
}

func (req SearchRequest) WithLanguages(languages ...string) *SearchRequest {
	req.Languages = languages
	return &req
}

func (req SearchRequest) WithAudios(audios ...string) *SearchRequest {
	req.Audios = audios
	return &req
}

func (req SearchRequest) WithSubtitles(subtitles ...string) *SearchRequest {
	req.Subtitles = subtitles
	return &req
}

func (req SearchRequest) WithSort(sort string) *SearchRequest {
	req.Sort = sort
	return &req
}

func (req SearchRequest) WithOrder(order string) *SearchRequest {
	req.Order = order
	return &req
}

func (req SearchRequest) WithPage(page int) *SearchRequest {
	req.Page = page
	return &req
}

// Do executes the search request against the client.
func (req SearchRequest) Do(ctx context.Context, cl *Client) (*SearchResponse, error) {
	res := new(SearchResponse)
	if err := cl.Do(ctx, "search", req, res); err != nil {
		return nil, err
	}
	switch {
	case res.StatusMessage != "":
		return nil, errors.New(res.StatusMessage)
	case !res.Success:
		return nil, errors.New("success != true")
	}
	return res, nil
}

// SearchResponse is a bhd search response.
type SearchResponse struct {
	// The status code of the post request. (0 = Failed and 1 = Success)
	StatusCode int `json:"status_code,omitempty"`
	// The current page of results that you're on.
	Page int `json:"page,omitempty"`
	// The results that match your query.
	Results []SearchResult `json:"results,omitempty"`
	// The total number of pages of results matching your query.
	TotalPages int `json:"total_pages,omitempty"`
	// The total number of results matching your query.
	TotalResults int `json:"total_results,omitempty"`
	// The status of the call. (True = Success, False = Error)
	Success bool `json:"success,omitempty"`
	// The status message.
	StatusMessage string `json:"status_message,omitempty"`
}

// SearchResult is a bhd search result.
type SearchResult struct {
	// The BHD ID.
	ID int `json:"id,omitempty"`
	// The name.
	Name string `json:"name,omitempty"`
	// The torrent folder name.
	FolderName string `json:"folder_name,omitempty"`
	// The torrent info_hash.
	InfoHash string `json:"info_hash,omitempty"`
	// The torrent size.
	Size int64 `json:"size,omitempty"`
	// The uploaders username.
	UploadedBy string `json:"uploaded_by,omitempty"`
	// The category.
	Category string `json:"category,omitempty"`
	// The type.
	Type string `json:"type,omitempty"`
	// The seeders.
	Seeders int `json:"seeders,omitempty"`
	// The leechers.
	Leechers int `json:"leechers,omitempty"`
	// The times completed.
	TimesCompleted int `json:"times_completed,omitempty"`
	// The ID of the matching IMDB page.
	ImdbID string `json:"imdb_id,omitempty"`
	// The ID of the matching TMDB page.
	TmdbID string `json:"tmdb_id,omitempty"`
	// The BHD rating.
	BhdRating float64 `json:"bhd_rating,omitempty"`
	// The TMDb rating.
	TmdbRating float64 `json:"tmdb_rating,omitempty"`
	// The IMDb rating.
	ImdbRating float64 `json:"imdb_rating,omitempty"`
	// The torrent TV pack flag.
	TvPack Bool `json:"tv_pack,omitempty"`
	// The torrent 25% promo.
	Promo25 Bool `json:"promo25,omitempty"`
	// The torrent 50% promo.
	Promo50 Bool `json:"promo50,omitempty"`
	// The torrent 75% promo.
	Promo75 Bool `json:"promo75,omitempty"`
	// The torrent freeleech status.
	Freeleech Bool `json:"freeleech,omitempty"`
	// The torrent rewind promo.
	Rewind Bool `json:"rewind,omitempty"`
	// The torrent refund promo.
	Refund Bool `json:"refund,omitempty"`
	// The torrent limited UL promo.
	Limited Bool `json:"limited,omitempty"`
	// The torrent rescue promo.
	Rescue Bool `json:"rescue,omitempty"`
	// The bumped at time.
	BumpedAt Time `json:"bumped_at,omitempty"`
	// The created at time.
	CreatedAt Time `json:"created_at,omitempty"`
	// The url.
	URL string `json:"url,omitempty"`
	// The download url.
	DownloadURL string `json:"download_url,omitempty"`
}

type Bool bool

func (b Bool) String() string {
	if b {
		return "1"
	}
	return "0"
}

func (b Bool) Int() int {
	if b {
		return 1
	}
	return 0
}

func (b Bool) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *Bool) UnmarshalJSON(buf []byte) error {
	switch string(bytes.ToLower(buf)) {
	case "true", "1":
		*b = true
		return nil
	case "false", "0":
		*b = false
		return nil
	}
	return fmt.Errorf("invalid bool value %q", buf)
}

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format(timefmt)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

func (t *Time) UnmarshalJSON(buf []byte) error {
	if len(buf) < 2 {
		return errors.New("invalid time value")
	}
	v, err := time.Parse(timefmt, string(buf[1:len(buf)-1]))
	if err != nil {
		return err
	}
	*t = Time(v)
	return nil
}

const timefmt = "2006-01-02 15:04:05"
