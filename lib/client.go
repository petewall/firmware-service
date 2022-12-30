package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	URL string
}

func (c *Client) get(path, name string) ([]byte, error) {
	res, err := http.Get(c.URL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to request the %s: %w", name, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request failed: %s", name, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the %s response: %w", name, err)
	}

	return body, nil
}

func (c *Client) GetAllFirmware() (FirmwareList, error) {
	name := "firmware list"
	body, err := c.get("/", name)
	if err != nil {
		return nil, err
	}

	var firmwareList FirmwareList
	err = json.Unmarshal(body, &firmwareList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return firmwareList, nil
}

func (c *Client) GetFirmwareTypes() ([]string, error) {
	name := "firmware types"
	body, err := c.get("/types", name)
	if err != nil {
		return nil, err
	}

	var types []string
	err = json.Unmarshal(body, &types)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return types, nil
}

func (c *Client) GetFirmwareByType(firmwareType string) (FirmwareList, error) {
	name := "firmware list for type " + firmwareType
	body, err := c.get("/"+firmwareType, name)
	if err != nil {
		return nil, err
	}

	var firmwareList FirmwareList
	err = json.Unmarshal(body, &firmwareList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return firmwareList, nil
}

func (c *Client) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	name := "firmware " + firmwareType + " " + firmwareVersion
	body, err := c.get("/"+firmwareType+"/"+firmwareVersion, name)
	if err != nil {
		return nil, err
	}

	var firmware *Firmware
	err = json.Unmarshal(body, &firmware)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return firmware, nil
}

func (c *Client) AddFirmware(firmwareType, firmwareVersion string, data []byte) error {
	req, err := http.NewRequest(http.MethodPut, c.URL+"/"+firmwareType+"/"+firmwareVersion, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create firmware %s %s upload request: %w", firmwareType, firmwareVersion, err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed request firmware %s %s upload: %w", firmwareType, firmwareVersion, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("firmware %s %s upload request failed: %s", firmwareType, firmwareVersion, res.Status)
	}

	return nil
}

func (c *Client) DeleteFirmware(firmwareType, firmwareVersion string) error {
	req, err := http.NewRequest(http.MethodDelete, c.URL+"/"+firmwareType+"/"+firmwareVersion, nil)
	if err != nil {
		return fmt.Errorf("failed to create firmware %s %s deletion request: %w", firmwareType, firmwareVersion, err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed request firmware %s %s deletion: %w", firmwareType, firmwareVersion, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("firmware %s %s deletion request failed: %s", firmwareType, firmwareVersion, res.Status)
	}

	return nil
}

func (c *Client) GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error) {
	name := "firmware " + firmwareType + " " + firmwareVersion + " data"
	return c.get("/"+firmwareType+"/"+firmwareVersion+"/data", name)
}
