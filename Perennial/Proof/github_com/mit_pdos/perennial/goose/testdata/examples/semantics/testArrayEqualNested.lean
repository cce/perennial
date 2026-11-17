/-
The proof of `testArrayEqualNested` (a goose semantics test, `arrays.go`).
-/
module

public import Perennial.Spec.github_com.mit_pdos.perennial.goose.testdata.examples.semantics.testArrayEqualNested

@[expose] public section

set_option linter.iris.style.nameCheck false
set_option linter.unusedSectionVars false

noncomputable section

namespace Perennial

open Iris Iris.BI Iris.ProgramLogic Iris.Std OFE

section proof
variable [GooseExts] [GooseSemantics] [GooseInterp]
variable [go_gctx : GoGlobalContext]
variable {hlc : HasLC} {GF : BundledGFunctors} [hG : HeapGS hlc GF]
variable [sem : go.Semantics]
variable [package_sem : github_com.mit_pdos.perennial.goose.testdata.examples.semantics.Assumptions]

theorem github_com.mit_pdos.perennial.goose.testdata.examples.semantics.testArrayEqualNested.wp_proof :
    github_com.mit_pdos.perennial.goose.testdata.examples.semantics.testArrayEqualNested.spec (GF := GF) (hG := hG) := by
  intro Φ
  iintro HΦ
  wp_func_call
  wp_call
  steps
  -- `==` on `[2]arrayInStruct` compares the elements with `go_eq_struct`: unfold the
  -- comparisons of the two indices
  simp only [show List.map (fun (i : Nat) => (i : Int)) (List.range 2) = [0, 1] from rfl,
    List.foldl_cons, List.foldl_nil]
  steps
  iexact HΦ

end proof

end Perennial
