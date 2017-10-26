package varis

type synapse struct {
	weight     float64
	uuid       string
	in         chan float64
	out        chan float64
	cache      float64
	inputUUID  string
	outputUUID string
}

func (syn *synapse) live() {
	for {
		syn.cache = <-syn.in
		outputValue := syn.cache * syn.weight
		syn.out <- outputValue
	}
}

func createSynapse(in Neuroner, out Neuroner, uuid string, weight float64) {
	syn := &synapse{
		weight:     weight,
		uuid:       uuid,
		in:         make(chan float64),
		out:        make(chan float64),
		inputUUID:  in.getUUID(),
		outputUUID: out.getUUID(),
	}

	in.getConnection().addOutputSynapse(syn)
	out.getConnection().addInputSynapse(syn)

	go syn.live()
}
