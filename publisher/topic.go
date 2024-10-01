package main

func createTopic(brokers []string, topic string) error {

	conn, err := kafka.Dail("tcp", brokers[0])

}
