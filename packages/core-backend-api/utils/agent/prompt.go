package agent

import (
	"fmt"
	"strings"
	"time"
)

type Prompt struct {
	ObjectiveAndPersona string
	Instructions        string
	Constraints         string
	Context             string
	OutputFormat        string
	Examples            string
}

func (p Prompt) objectiveAndPersonaPrompt() string {
	if p.ObjectiveAndPersona == "" {
		return ""
	}

	t := `
    <OBJECTIVE_AND_PERSONA>
        %s
    </OBJECTIVE_AND_PERSONA>
    `

	return fmt.Sprintf(t, p.ObjectiveAndPersona)
}

func (p Prompt) instructionsPrompt() string {
	if p.Instructions == "" {
		return ""
	}

	t := `
    <INSTRUCTIONS>
        %s
    </INSTRUCTIONS>
    `

	return fmt.Sprintf(t, p.Instructions)
}

func (p Prompt) constraintsPrompt() string {
	if p.Constraints == "" {
		return ""
	}

	t := `
    <CONSTRAINTS>
        %s
    </CONSTRAINTS>
    `

	return fmt.Sprintf(t, p.Constraints)
}

func (p Prompt) contextPrompt() string {
	now := time.Now()
	today := now.Format("26th January, 2025")

	t := `
    <CONTEXT>
		Today is %s
		User Language is English
		User is currently in Berlin, Germany
		User's currency is EURO
        %s
    </CONTEXT>
    `

	return fmt.Sprintf(t, today, p.Context)
}

func (p Prompt) outputFormatPrompt() string {
	if p.OutputFormat == "" {
		return ""
	}

	t := `
    <OUTPUT_FORMAT>
        %s
    </OUTPUT_FORMAT>
    `

	return fmt.Sprintf(t, p.OutputFormat)
}

func (p Prompt) examplesPrompt() string {
	if p.Examples == "" {
		return ""
	}

	t := `
    <FEW_SHOT_EXAMPLES>
        %s
    </FEW_SHOT_EXAMPLES>
    `

	return fmt.Sprintf(t, p.Examples)
}

func (p Prompt) Stitch() string {
	return strings.Join([]string{
		p.objectiveAndPersonaPrompt(),
		p.instructionsPrompt(),
		p.constraintsPrompt(),
		p.contextPrompt(),
		p.outputFormatPrompt(),
		p.examplesPrompt(),
	}, "\n\n")
}
