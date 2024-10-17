package utils

type StepFunc func(interface{}) error

func Run(steps []StepFunc, intention interface{}) error {
	for _, step := range steps {
		if err := step(intention); err != nil {
			return err
		}
	}
	return nil
}
