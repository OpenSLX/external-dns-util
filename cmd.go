package main

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/registry"
)

func execCmd(ctx context.Context, r registry.Registry) {
	cfg := externaldns.GlobalCmdArgs

	record := endpoint.NewEndpoint(cfg.CommandName, cfg.CommandType, cfg.CommandTarget)
	if cfg.CommandTTL != 0 {
		record = endpoint.NewEndpointWithTTL(cfg.CommandName, cfg.CommandType,
			endpoint.TTL(cfg.CommandTTL), cfg.CommandTarget)
	}

	switch cfg.Command {
	case "show":
		recs, err := r.Records(ctx)
		if err != nil {
			log.Fatal(err)
		}
		json, err := json.Marshal(recs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json))
		break
	case "add":
		err := r.ApplyChanges(ctx, &plan.Changes{Create: []*endpoint.Endpoint{record}})
		if err != nil {
			log.Fatal(err)
		}
		break
	case "del":
		recs, err := r.Records(ctx)
		if err != nil {
			log.Fatal(err)
		}
		recsDel := []*endpoint.Endpoint{}
		for _, v := range recs {
			if v.DNSName == record.DNSName && v.RecordType == v.RecordType {
				recsDel = append(recsDel, v)
			}
		}
		err = r.ApplyChanges(ctx, &plan.Changes{Delete: recsDel})
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
