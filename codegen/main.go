package main

import "fmt"
import "reflect"

type GlobalKind struct {
	val int
}

func Transfer() GlobalKind {
	return GlobalKind{val: 0}
}
func Branch() GlobalKind {
	return GlobalKind{val: 1}
}
func GlobalLoop() GlobalKind {
	return GlobalKind{val: 2}
}
func GlobalBreak() GlobalKind {
	return GlobalKind{val: 3}
}

type IGlobalAction interface {
	isGlobalAction()
}

type GlobalAction struct {
	kind   GlobalKind
	action IGlobalAction
}

func (a *GlobalAction) print(indent int) {
	switch a.kind {
	case Transfer():
		a := a.action.(TransferAction)
		a.print(indent)
	case Branch():
		a := a.action.(BranchAction)
		a.print(indent)
	case GlobalLoop():
		a := a.action.(GlobalLoopAction)
		a.print(indent)
	case GlobalBreak():
		a.action.(GlobalBreakAction).print(indent)
	default:
		panic("Unknown global action when printing")
	}
}

type Field struct {
	ty   string
	name string
}

type Message struct {
	fields []Field
}

type TransferAction struct {
	sender      string
	receiver    string
	messageType string
}

func (a TransferAction) isGlobalAction() {}

func transferAction(action TransferAction) GlobalAction {
	return GlobalAction{
		kind:   Transfer(),
		action: action,
	}
}

func (a *TransferAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("%s -> %s: %s;\n", a.sender, a.receiver, a.messageType)
}

type BranchOption struct {
	messageType string
	actions     []GlobalAction
}

func (b *BranchOption) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println(b.messageType, "{")
	for _, action := range b.actions {
		action.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type BranchAction struct {
	sender   string
	receiver string
	branches []BranchOption
}

func (a BranchAction) isGlobalAction() {}

func branchAction(action BranchAction) GlobalAction {
	return GlobalAction{
		kind:   Branch(),
		action: action,
	}
}

func (a *BranchAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("%s -> %s {\n", a.sender, a.receiver)
	for _, branch := range a.branches {
		branch.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type GlobalLoopAction struct {
	actions []GlobalAction
}

func (a GlobalLoopAction) isGlobalAction() {}

func globalLoopAction(action GlobalLoopAction) GlobalAction {
	return GlobalAction{
		kind:   GlobalLoop(),
		action: action,
	}
}

func (a *GlobalLoopAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("loop {")
	for _, action := range a.actions {
		action.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type GlobalBreakAction struct{}

func (a GlobalBreakAction) isGlobalAction() {}

func globalBreakAction() GlobalAction {
	return GlobalAction{
		kind:   GlobalBreak(),
		action: GlobalBreakAction{},
	}
}

func (a GlobalBreakAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("break")
}

type GlobalProtocol struct {
	roles   []string
	actions []GlobalAction
}

func (p *GlobalProtocol) print() {
	fmt.Println("ROLES")
	for _, role := range p.roles {
		fmt.Println(role)
	}
	fmt.Println("\nPROTOCOL")
	for _, action := range p.actions {
		action.print(0)
	}
}

type LocalKind struct {
	val int
}

func Send() LocalKind {
	return LocalKind{val: 0}
}
func Receive() LocalKind {
	return LocalKind{val: 1}
}
func Choose() LocalKind {
	return LocalKind{val: 2}
}
func Offer() LocalKind {
	return LocalKind{val: 3}
}
func LocalLoop() LocalKind {
	return LocalKind{val: 4}
}
func LocalBreak() LocalKind {
	return LocalKind{val: 5}
}

type ILocalAction interface {
	isLocalAction()
}

type LocalAction struct {
	kind   LocalKind
	action ILocalAction
}

func (n *LocalAction) print(indent int) {
	switch n.kind {
	case Send():
		n := n.action.(SendAction)
		n.print(indent)
	case Receive():
		n := n.action.(ReceiveAction)
		n.print(indent)
	case Choose():
		n := n.action.(ChooseAction)
		n.print(indent)
	case Offer():
		n := n.action.(OfferAction)
		n.print(indent)
	case LocalLoop():
		n := n.action.(LocalLoopAction)
		n.print(indent)
	case LocalBreak():
		n.action.(LocalBreakAction).print(indent)
	default:
		panic("Unknown local action when printing")
	}
}

type SendAction struct {
	receiver    string
	messageType string
}

func (a SendAction) isLocalAction() {}

func sendAction(receiver string, messageType string) LocalAction {
	return LocalAction{
		kind:   Send(),
		action: SendAction{receiver: receiver, messageType: messageType},
	}
}

func (a *SendAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("!%s: %s\n", a.receiver, a.messageType)
}

type ReceiveAction struct {
	sender      string
	messageType string
}

func (a ReceiveAction) isLocalAction() {}

func receiveAction(sender string, messageType string) LocalAction {
	return LocalAction{
		kind:   Receive(),
		action: ReceiveAction{sender: sender, messageType: messageType},
	}
}

func (a *ReceiveAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("?%s: %s\n", a.sender, a.messageType)
}

type ChooseOption struct {
	messageType  string
	continuation []LocalAction
}

func (c *ChooseOption) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println(c.messageType, "{")
	for _, action := range c.continuation {
		action.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type ChooseAction struct {
	receiver string
	choices  []ChooseOption
}

func (a ChooseAction) isLocalAction() {}

func chooseAction(receiver string, choices []ChooseOption) LocalAction {
	return LocalAction{
		kind:   Choose(),
		action: ChooseAction{receiver: receiver, choices: choices},
	}
}

func (a *ChooseAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("!%s {\n", a.receiver)
	for _, choice := range a.choices {
		choice.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type OfferOption struct {
	messageType  string
	continuation []LocalAction
}

func (c *OfferOption) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println(c.messageType, "{")
	for _, action := range c.continuation {
		action.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type OfferAction struct {
	sender string
	offers []OfferOption
}

func (a OfferAction) isLocalAction() {}

func offerAction(sender string, offers []OfferOption) LocalAction {
	return LocalAction{
		kind:   Offer(),
		action: OfferAction{sender: sender, offers: offers},
	}
}

func (a *OfferAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Printf("?%s {\n", a.sender)
	for _, offer := range a.offers {
		offer.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type LocalLoopAction struct {
	actions []LocalAction
}

func (a LocalLoopAction) isLocalAction() {}

func localLoopAction(actions []LocalAction) LocalAction {
	return LocalAction{
		kind:   LocalLoop(),
		action: LocalLoopAction{actions: actions},
	}
}

func (a LocalLoopAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("loop {")
	for _, action := range a.actions {
		action.print(indent + 1)
	}
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("}")
}

type LocalBreakAction struct{}

func (a LocalBreakAction) isLocalAction() {}

func localBreakAction() LocalAction {
	return LocalAction{
		kind:   LocalBreak(),
		action: LocalBreakAction{},
	}
}

func (a LocalBreakAction) print(indent int) {
	for range indent {
		fmt.Print("\t")
	}
	fmt.Println("break")
}

type LocalProtocol struct {
	actions []LocalAction
}

func (p *LocalProtocol) print() {
	for _, action := range p.actions {
		action.print(0)
	}
}

func mergeBranches(branches [][]LocalAction) (string, []OfferOption, error) {
	sender := ""
	offers := make([]OfferOption, 0)
	uniqueBranches := make(map[string][]LocalAction, 0)
	hasEmptyBranch := false

	for i, branch := range branches {
		if hasEmptyBranch && len(branch) != 0 {
			return "", nil, fmt.Errorf("all branches must be empty or non-empty")
		} else if !hasEmptyBranch && len(branch) == 0 && i == len(branches) {
			return "", nil, fmt.Errorf("all branches must be empty or non-empty")
		} else if len(branch) == 0 {
			hasEmptyBranch = true
			continue
		}

		var recv *ReceiveAction = nil
		switch branch[0].kind {
		case Receive():
			a := branch[0].action.(ReceiveAction)
			recv = &a
		default:
			return "", nil, fmt.Errorf("implicit offer branches must begin with a receive")
		}

		if recv.sender == "" {
			panic("verify sender is not empty before merging branches")
		} else if recv.messageType == "" {
			panic("verify message type is not empty before merging branches")
		}

		if sender == "" {
			sender = recv.sender
		} else if sender != recv.sender {
			return "", nil, fmt.Errorf("all branches must have the same sender")
		}

		if b, ok := uniqueBranches[recv.messageType]; ok {
			if reflect.DeepEqual(b, branch[1:]) {
				continue
			} else {
				return "", nil, fmt.Errorf("all branches must either be unique or have a distinguished receive type")
			}
		}
		uniqueBranches[recv.messageType] = branch[1:]
		offers = append(offers, OfferOption{messageType: recv.messageType, continuation: branch[1:]})
	}
	return sender, offers, nil
}

func recursivelyProject(globalActions []GlobalAction, localActions map[string][]LocalAction) error {
	if len(localActions) == 0 {
		panic("local actions must be initialized to empty arrays")
	}

	for _, action := range globalActions {
		switch action.kind {
		case Transfer():
			a := action.action.(TransferAction)
			// TODO: Check if sender, receiver, and message type are not empty
			if actions, ok := localActions[a.sender]; ok {
				localActions[a.sender] = append(actions, sendAction(a.receiver, a.messageType))
			} else {
				return fmt.Errorf("sender role `%s` does not exist", a.sender)
			}
			if actions, ok := localActions[a.receiver]; ok {
				localActions[a.receiver] = append(actions, receiveAction(a.sender, a.messageType))
			} else {
				return fmt.Errorf("receiver role `%s` does not exist", a.receiver)
			}
		case Branch():
			a := action.action.(BranchAction)
			choices := make([]ChooseOption, 0)
			offers := make([]OfferOption, 0)
			branchBundles := make(map[string][][]LocalAction)
			for _, globalBranch := range a.branches {
				localBranches := make(map[string][]LocalAction)
				for role := range localActions {
					localBranches[role] = make([]LocalAction, 0)
				}
				recursivelyProject(globalBranch.actions, localBranches)
				choices = append(choices, ChooseOption{messageType: globalBranch.messageType, continuation: localBranches[a.sender]})
				offers = append(offers, OfferOption{messageType: globalBranch.messageType, continuation: localBranches[a.receiver]})
				for role, branch := range localBranches {
					if role == a.sender || role == a.receiver {
						continue
					}
					branchBundles[role] = append(branchBundles[role], branch)
				}
			}
			localActions[a.sender] = append(localActions[a.sender], chooseAction(a.receiver, choices))
			localActions[a.receiver] = append(localActions[a.receiver], offerAction(a.sender, offers))
			for role, branches := range branchBundles {
				sender, offers, err := mergeBranches(branches)
				if err != nil {
					return err
				}
				if len(offers) != 0 {
					localActions[role] = append(localActions[role], offerAction(sender, offers))
				}
			}
		case GlobalLoop():
			a := action.action.(GlobalLoopAction)
			continuations := make(map[string][]LocalAction)
			for role := range localActions {
				continuations[role] = make([]LocalAction, 0)
			}
			err := recursivelyProject(a.actions, continuations)
			if err != nil {
				return err
			}
			for role, actions := range continuations {
				if len(actions) == 0 || (len(actions) == 1 && actions[0].kind == LocalBreak()) {
					// NOTE: If a role does not appear in the loop, then
					// the loop may as well not exist for that role
				} else {
					localActions[role] = append(localActions[role], localLoopAction(actions))
				}
			}
		case GlobalBreak():
			for role, actions := range localActions {
				localActions[role] = append(actions, localBreakAction())
			}
		default:
			panic("Unknown global action when projecting to local")
		}
	}
	return nil
}

func projectGlobalToLocal(global GlobalProtocol) (map[string]LocalProtocol, error) {
	localActions := make(map[string][]LocalAction)
	for _, role := range global.roles {
		localActions[role] = make([]LocalAction, 0)
	}

	err := recursivelyProject(global.actions, localActions)
	if err != nil {
		return nil, err
	}

	p := make(map[string]LocalProtocol)
	for role, actions := range localActions {
		p[role] = LocalProtocol{actions: actions}
	}
	return p, nil
}

func main() {
	globalProtocol := GlobalProtocol{
		roles: []string{"buyer1", "buyer2", "seller"},
		actions: []GlobalAction{
			transferAction(TransferAction{
				sender:      "buyer1",
				receiver:    "seller",
				messageType: "title",
			}),
			transferAction(TransferAction{
				sender:      "seller",
				receiver:    "buyer1",
				messageType: "price",
			}),
			transferAction(TransferAction{
				sender:      "buyer1",
				receiver:    "buyer2",
				messageType: "split",
			}),
			globalLoopAction(GlobalLoopAction{
				actions: []GlobalAction{
					branchAction(BranchAction{
						sender:   "buyer2",
						receiver: "buyer1",
						branches: []BranchOption{
							{messageType: "accept", actions: []GlobalAction{
								transferAction(TransferAction{sender: "buyer1", receiver: "seller", messageType: "done"}),
								globalBreakAction(),
							}},
							{messageType: "quit", actions: []GlobalAction{
								transferAction(TransferAction{sender: "buyer1", receiver: "seller", messageType: "done"}),
								globalBreakAction(),
							}},
							{messageType: "retry", actions: []GlobalAction{
								transferAction(TransferAction{sender: "buyer1", receiver: "seller", messageType: "retry"}),
							}},
						},
					}),
				},
			}),
		},
	}

	localProtocols, err := projectGlobalToLocal(globalProtocol)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	globalProtocol.print()

	for role, localProtocol := range localProtocols {
		fmt.Println("=====================================")
		fmt.Println(role)
		fmt.Println("=====================================")
		localProtocol.print()
		fmt.Println()
	}
}
