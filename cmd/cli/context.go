package main

import (
	"fmt"
	"os"
	"time"

	"github.com/javuto/osctrl/context"
	"github.com/urfave/cli"
)

func addContext(c *cli.Context) error {
	// Get context name
	ctxName := c.String("name")
	if ctxName == "" {
		fmt.Println("Context name is required")
		os.Exit(1)
	}
	// Get context hostname
	ctxHost := c.String("hostname")
	if ctxHost == "" {
		fmt.Println("Context hostname is required")
		os.Exit(1)
	}
	// Get configuration
	var configuration string
	confFile := c.String("configuration")
	if confFile != "" {
		configuration = context.ReadExternalFile(confFile)
	}
	// Get certificate
	var certificate string
	certFile := c.String("certificate")
	if certFile != "" {
		certificate = context.ReadExternalFile(certFile)
	}
	// Create context if it does not exist
	if !ctxs.Exists(ctxName) {
		newContext := ctxs.Empty(ctxName, ctxHost)
		newContext.DebugHTTP = c.Bool("debug")
		newContext.Configuration = configuration
		newContext.Certificate = certificate
		newContext.EnrollExpire = time.Now().Add(time.Duration(context.DefaultLinkExpire) * time.Hour)
		newContext.RemoveExpire = time.Now().Add(time.Duration(context.DefaultLinkExpire) * time.Hour)
		if err := ctxs.Create(newContext); err != nil {
			return err
		}
	} else {
		fmt.Printf("Context %s already exists!\n", ctxName)
		os.Exit(1)
	}
	return nil
}

func deleteContext(c *cli.Context) error {
	// Get context name
	ctxName := c.String("name")
	if ctxName == "" {
		fmt.Println("Context name is required")
		os.Exit(1)
	}
	return ctxs.Delete(ctxName)
}

func showContext(c *cli.Context) error {
	// Get context name
	ctxName := c.String("name")
	if ctxName == "" {
		fmt.Println("Context name is required")
		os.Exit(1)
	}
	ctx, err := ctxs.Get(ctxName)
	if err != nil {
		return err
	}
	fmt.Printf("Context %s\n", ctx.Name)
	fmt.Printf(" Name: %s\n", ctx.Name)
	fmt.Printf(" Host: %s\n", ctx.Hostname)
	fmt.Printf(" Secret: %s\n", ctx.Secret)
	fmt.Printf(" EnrollExpire: %v\n", ctx.EnrollExpire)
	fmt.Printf(" EnrollSecretPath: %s\n", ctx.EnrollSecretPath)
	fmt.Printf(" RemoveExpire: %v\n", ctx.RemoveExpire)
	fmt.Printf(" RemoveSecretPath: %s\n", ctx.RemoveSecretPath)
	fmt.Printf(" Type: %v\n", ctx.Type)
	fmt.Printf(" DebugHTTP? %v\n", ctx.DebugHTTP)
	fmt.Printf(" Icon: %s\n", ctx.Icon)
	fmt.Println(" Configuration: ")
	fmt.Printf("%s\n", ctx.Configuration)
	fmt.Println(" Certificate: ")
	fmt.Printf("%s\n", ctx.Certificate)
	fmt.Println()
	return nil
}

func listContext(c *cli.Context) error {
	contexts, err := ctxs.All()
	if err != nil {
		return err
	}
	if len(contexts) > 0 {
		fmt.Printf("Existing contexts:\n\n")
		for _, ctx := range contexts {
			fmt.Printf("  Name: %s\n", ctx.Name)
		}
		fmt.Println()
	} else {
		fmt.Printf("No contexts\n")
	}
	return nil
}

func quickAddContext(c *cli.Context) error {
	// Get context name
	ctxName := c.String("name")
	if ctxName == "" {
		fmt.Println("Context name is required")
		os.Exit(1)
	}
	ctx, err := ctxs.Get(ctxName)
	if err != nil {
		return err
	}
	var oneLiner string
	if c.String("target") == "sh" {
		oneLiner, _ = context.QuickAddOneLinerShell(ctx)
	} else if c.String("target") == "ps1" {
		oneLiner, _ = context.QuickAddOneLinerPowershell(ctx)
	}
	fmt.Printf("%s\n", oneLiner)
	return nil
}