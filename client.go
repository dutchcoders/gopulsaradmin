package pulsaradmin

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Key string

	debug bool
	*http.Client
	baseURL *url.URL
}

type TopicStats struct {
	AverageMsgSize         float64 `json:"averageMsgSize"`
	MsgRateIn              float64 `json:"msgRateIn"`
	MsgRateOut             float64 `json:"msgRateOut"`
	MsgThroughputIn        float64 `json:"msgThroughputIn"`
	MsgThroughputOut       float64 `json:"msgThroughputOut"`
	PendingAddEntriesCount int64   `json:"pendingAddEntriesCount"`
	ProducerCount          int64   `json:"producerCount"`
	Publishers             []struct {
		Address        string  `json:"address"`
		AverageMsgSize float64 `json:"averageMsgSize"`
		ClientVersion  string  `json:"clientVersion"`
		ConnectedSince string  `json:"connectedSince"`
		Metadata       struct {
		} `json:"metadata"`
		MsgRateIn       float64 `json:"msgRateIn"`
		MsgThroughputIn float64 `json:"msgThroughputIn"`
		ProducerId      int64   `json:"producerId"`
		ProducerName    string  `json:"producerName"`
	} `json:"publishers"`
	Replication struct {
	} `json:"replication"`
	StorageSize   int64 `json:"storageSize"`
	Subscriptions struct {
		Subscription struct {
			BlockedSubscriptionOnUnackedMsgs bool `json:"blockedSubscriptionOnUnackedMsgs"`
			Consumers                        []struct {
				Address                      string `json:"address"`
				AvailablePermits             int64  `json:"availablePermits"`
				BlockedConsumerOnUnackedMsgs bool   `json:"blockedConsumerOnUnackedMsgs"`
				ClientVersion                string `json:"clientVersion"`
				ConnectedSince               string `json:"connectedSince"`
				ConsumerName                 string `json:"consumerName"`
				Metadata                     struct {
				} `json:"metadata"`
				MsgRateOut       float64 `json:"msgRateOut"`
				MsgRateRedeliver float64 `json:"msgRateRedeliver"`
				MsgThroughputOut float64 `json:"msgThroughputOut"`
				UnackedMessages  int64   `json:"unackedMessages"`
			} `json:"consumers"`
			MsgBacklog                               int64   `json:"msgBacklog"`
			MsgRateExpired                           float64 `json:"msgRateExpired"`
			MsgRateOut                               float64 `json:"msgRateOut"`
			MsgRateRedeliver                         float64 `json:"msgRateRedeliver"`
			MsgThroughputOut                         float64 `json:"msgThroughputOut"`
			NumberOfEntriesSinceFirstNotAckedMessage int64   `json:"numberOfEntriesSinceFirstNotAckedMessage"`
			TotalNonContiguousDeletedMessagesRange   int64   `json:"totalNonContiguousDeletedMessagesRange"`
			Type                                     string  `json:"type"`
			UnackedMessages                          int64   `json:"unackedMessages"`
		} `json:"subscription"`
	} `json:"subscriptions"`
}

type BrokerStatsTopicResult map[string]Namespace

type Namespace map[string]NamespaceBundle

type NamespaceBundle map[string]TopicType

type TopicType map[string]TopicStats

func (c *Client) BrokerStatsTopics() (BrokerStatsTopicResult, error) {
	req, err := c.NewRequest("GET", "/admin/v2/broker-stats/topics", nil)
	if err != nil {
		return nil, err
	}

	resp := BrokerStatsTopicResult{}
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type optionFn func(*Client)

// WithDebug enables debug output while interacting with MIPS
func WithDebug() optionFn {
	return func(c *Client) {
		c.debug = true
	}
}

// WithURL contains the MIPS target url
func WithURL(u url.URL) optionFn {
	return func(c *Client) {
		c.baseURL = &u
	}
}

// WithKey contains the MIPS API key
func WithKey(key string) optionFn {
	return func(c *Client) {
		c.Key = key
	}
}

// New returns a MIPS API client
func New(options ...optionFn) (*Client, error) {
	c := &Client{
		Client: http.DefaultClient,
	}

	for _, optionFn := range options {
		optionFn(c)
	}

	if c.baseURL == nil {
		return nil, fmt.Errorf("URL not set")
	}

	return c, nil
}
