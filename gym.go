package gym

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const dataFile = "members.json"

type Member struct {
	MemberID        string `json:"member_id"`
	FullName        string `json:"full_name"`
	PlanType        string `json:"plan_type"`
	MonthsRemaining int    `json:"months_remaining"`
	IsActive        bool   `json:"is_active"`
}

func loadMembers() (map[string]Member, error) {
	data, err := os.ReadFile(dataFile)
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]Member), nil
	}
	if err != nil {
		return nil, err
	}

	members := make(map[string]Member)
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, fmt.Errorf("could not read %s: %w", dataFile, err)
	}
	return members, nil
}

func saveMembers(members map[string]Member) error {
	data, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func EnrollMember(member Member) error {
	member.MemberID = strings.TrimSpace(member.MemberID)
	member.FullName = strings.TrimSpace(member.FullName)
	member.PlanType = strings.TrimSpace(member.PlanType)

	if member.MemberID == "" {
		return errors.New("member ID cannot be empty")
	}
	if member.FullName == "" {
		return errors.New("full name cannot be empty")
	}
	if member.PlanType != "Standard" && member.PlanType != "Premium" && member.PlanType != "VIP" {
		return errors.New("invalid plan; choose Standard, Premium, or VIP")
	}
	if member.MonthsRemaining <= 0 {
		return errors.New("months must be greater than zero")
	}

	members, err := loadMembers()
	if err != nil {
		return err
	}
	if _, exists := members[member.MemberID]; exists {
		return fmt.Errorf("member ID %q already exists", member.MemberID)
	}

	members[member.MemberID] = member
	return saveMembers(members)
}

func LookupMember(id string) (*Member, error) {
	members, err := loadMembers()
	if err != nil {
		return nil, err
	}

	member, exists := members[strings.TrimSpace(id)]
	if !exists {
		return nil, errors.New("member not found")
	}
	return &member, nil
}

func RenewSubscription(id string, extraMonths int) error {
	if extraMonths <= 0 {
		return errors.New("months must be greater than zero")
	}

	members, err := loadMembers()
	if err != nil {
		return err
	}

	id = strings.TrimSpace(id)
	member, exists := members[id]
	if !exists {
		return errors.New("member not found")
	}
	member.MonthsRemaining += extraMonths
	member.IsActive = true
	members[id] = member
	return saveMembers(members)
}

func RevokeMembership(id string) error {
	members, err := loadMembers()
	if err != nil {
		return err
	}

	id = strings.TrimSpace(id)
	member, exists := members[id]
	if !exists {
		return errors.New("member not found")
	}
	member.IsActive = false
	members[id] = member
	return saveMembers(members)
}
