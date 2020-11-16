(defrule addtoanswer
  ?fact <- (answer ?theanswer)
  =>
  (assert (answser (+ 1 ?theanswer)))
  (retract ?fact)
)
